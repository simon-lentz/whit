// Package scrapers contains various data extraction tools.
package scrapers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// Facility models base object.
type Facility struct {
	State          string `csv:"state" json:"state"`
	Name           string `csv:"name" json:"name"`
	OwnerPERMITTEE string `csv:"owner" json:"owner"`
	ProjectType    string `csv:"project_type" json:"project_type"`
	Capacity       string `csv:"capacity_mw" json:"capacity_mw"`
	PSCapacity     string `csv:"ps_capacity_mw" json:"ps_capacity_mw"`
	Waterway       string `csv:"waterway" json:"waterway"`
}

// Output models a slice of Facility structs, facilitates uniform json output.
type Output struct {
	Facilities []Facility `json:"facilities"`
}

var states = []string{"AL", "AR", "AZ", "CA", "CO", "CT", "DE", "DC", "FL",
	"GA", "ID", "IN", "IA", "KS", "KY", "LA", "ME", "MD",
	"MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ",
	"NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC",
	"SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}
var header = []string{"state", "name", "owner_PERMITTEE", "project_type",
	"capacity_mw", "ps_capacity_mw", "waterway"}
var targets = []string{exist, pump, pipe}

var filename string
var outdir string
var fileType string

// NhaCmd is a command line user interface for scraping the NHA database.
var NhaCmd = &cobra.Command{
	Use:   "nha",
	Short: "scrapes National Hydropower Association (nha) database",
	Long: `
The subcommand <nha> scrapes the National Hydropower Association
output filename must be specified "whit collect nha -f filename" 
output path defaults to $HOME but can be specified with using -o,
e.g. "whit collect nha -f filename -o /dir/dir2/targetdir"`,
	Run: func(cmd *cobra.Command, args []string) {
		switch fT := fileType; fT {
		case "json":
			log.Println("Scraping NHA database...")
			writeJSON(filename, outdir)
			log.Printf("Scrape complete, output data written to %s.json", filename)
		default:
			log.Println("Scraping NHA database...")
			writeCSV(filename, outdir)
			log.Printf("Scrape complete, output data written to %s.csv", filename)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

const (
	exist = "existing-hydropower"
	pipe  = "hydropower-pipeline"
	pump  = "pumped-storage"
	null  = "null"
)

// scraper iterates by over the target url by state, extracts and transforms data, then returns an array of normalized string data.
func scraper(state, target string) [][]string {
	var facData [][]string
	c := colly.NewCollector()
	// c.OnHTML callback starts scraping where div[class="facility-details"].
	c.OnHTML(".facility-details", func(e *colly.HTMLElement) {
		// element is the goquery DOM.
		element := e.DOM
		// childNodes iterates through DOM targets children.
		childNodes := element.Children().Nodes
		// custom text transformation with strings.NewReplacer().
		r := strings.NewReplacer("&", "and", ",", "", ": ", "", "Ownership", "", "Owner", "", "Permittee", "", "Type", "",
			"Capacity", "", "Authorized ", "", " (MW)", "", "Hydropower Megawatts", "",
			"Pumped Storage Megawatts", "", " of Permit", "", "Waterway", "", "\"", "", "\n", "")
		switch fac := target; fac {
		case exist:
			existFac := Facility{
				State:          state,
				Name:           r.Replace(element.FindNodes(childNodes[0]).Text()),
				OwnerPERMITTEE: r.Replace(element.FindNodes(childNodes[1]).Text()),
				ProjectType:    r.Replace(element.FindNodes(childNodes[2]).Text()),
				Capacity:       r.Replace(element.FindNodes(childNodes[3]).Text()),
				PSCapacity:     null,
				Waterway:       null,
			}
			facData = append(facData, []string{existFac.State, existFac.Name, existFac.OwnerPERMITTEE,
				existFac.ProjectType, existFac.Capacity, existFac.PSCapacity, existFac.Waterway})
		case pipe:
			pipeFac := Facility{
				State:          state,
				Name:           r.Replace(element.FindNodes(childNodes[0]).Text()),
				OwnerPERMITTEE: r.Replace(element.FindNodes(childNodes[1]).Text()),
				ProjectType:    r.Replace(element.FindNodes(childNodes[2]).Text()),
				Capacity:       r.Replace(element.FindNodes(childNodes[3]).Text()),
				PSCapacity:     null,
				Waterway:       r.Replace(element.FindNodes(childNodes[4]).Text()),
			}
			facData = append(facData, []string{pipeFac.State, pipeFac.Name, pipeFac.OwnerPERMITTEE,
				pipeFac.ProjectType, pipeFac.Capacity, pipeFac.PSCapacity, pipeFac.Waterway})
		case pump:
			psFac := Facility{
				State:          state,
				Name:           r.Replace(element.FindNodes(childNodes[0]).Text()),
				OwnerPERMITTEE: r.Replace(element.FindNodes(childNodes[1]).Text()),
				ProjectType:    r.Replace(element.FindNodes(childNodes[2]).Text()),
				Capacity:       r.Replace(element.FindNodes(childNodes[3]).Text()),
				PSCapacity:     r.Replace(element.FindNodes(childNodes[4]).Text()),
				Waterway:       null,
			}
			facData = append(facData, []string{psFac.State, psFac.Name, psFac.OwnerPERMITTEE,
				psFac.ProjectType, psFac.Capacity, psFac.PSCapacity, psFac.Waterway})
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("URL", err)
	})
	if err := c.Visit(fmt.Sprintf("https://hydro.org/map/hydro/%s/?state=%s", target, state)); err != nil {
		log.Println(err)
	}
	return facData
}

// writeCSV creates and writes to a csv file (default output option).
func writeCSV(filename, outdir string) {
	switch out := outdir; out {
	case "":
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Println(err)
		}
		file, err := os.Create(homedir + "/" + filename + ".csv") //nolint:gosec
		if err != nil {
			log.Println(err)
		}
		w := csv.NewWriter(file)
		if err := w.Write(header); err != nil {
			log.Println(err)
		}
		for _, target := range targets {
			for _, state := range states {
				if err := w.WriteAll(scraper(state, target)); err != nil {
					log.Println(err)
				}
			}
		}
		w.Flush()
	default:
		file, err := os.Create(outdir + "/" + filename + ".csv") //nolint:gosec
		if err != nil {
			log.Println("error creating file:", err)
		}
		w := csv.NewWriter(file)
		if err := w.Write(header); err != nil {
			log.Println(err)
		}
		for _, target := range targets {
			for _, state := range states {
				if err := w.WriteAll(scraper(state, target)); err != nil {
					log.Println(err)
				}
			}
		}
		w.Flush()
	}
}

// writeJSON creates and writes to a json file (triggered by -t flag).
func writeJSON(filename, outdir string) {
	output := new(Output)
	for _, target := range targets {
		for _, state := range states {
			data := scraper(state, target)
			for i := 0; i < len(data); i++ {
				fac := Facility{
					State:          data[i][0],
					Name:           data[i][1],
					OwnerPERMITTEE: data[i][2],
					ProjectType:    data[i][3],
					Capacity:       data[i][4],
					PSCapacity:     data[i][5],
					Waterway:       data[i][6],
				}
				output.Facilities = append(output.Facilities, fac)
			}
		}
		switch out := outdir; out {
		case "":
			homedir, err := os.UserHomeDir()
			if err != nil {
				log.Println(err)
			}
			file, err := os.Create(homedir + "/" + filename + ".json") //nolint:gosec
			if err != nil {
				log.Println(err)
			}
			w := json.NewEncoder(file)
			w.SetIndent("", "  ")
			if err := w.Encode(output); err != nil {
				log.Println(err)
			}
		default:
			file, err := os.Create(outdir + "/" + filename + ".json") //nolint:gosec
			if err != nil {
				log.Println(err)
			}
			w := json.NewEncoder(file)
			w.SetIndent("", "  ")
			if err := w.Encode(output); err != nil {
				log.Println(err)
			}
		}
	}
}

func init() {
	flags := NhaCmd.Flags()
	flags.StringVarP(&filename, "filename", "f", "", "set output file name, required")
	if err := NhaCmd.MarkFlagRequired("filename"); err != nil {
		log.Println(err)
	}
	flags.StringVarP(&outdir, "outdir", "o", "", "set $PATH for output file, defaults to $HOME")
	flags.StringVarP(&fileType, "filetype", "t", "", "change output type to `json`, default output is csv")
}

/*
To Do:

+ The scraper function can probably be rewritten to use one Facility
  struct that extracts a broader range of target data, rather than
  the narrow switch statement cases. This will address a few major issues:
  1. The scraper is fragile. If the target HTML is modified in any way, the
     scraper breaks. A more rigorous implementation would prevent this.
  2. The use of narrow cases limits the reproducability of the scraper
     function across different sites. A more robust implementation
     would allow for this original scraper to be used as a more effective
     template for the many additional scraper needed.

+ With larger targets the chances of getting IP-blocked from target site
  are higher. This has not been an issue with the nha site, so I have not
  yet implemented any workarounds.
  1. Decide on where in the structure of `whit` this should be implemented.
  2. Decide on the method of workaround, i.e. how many degrees of separation
     should exist between request source and server target.
  3. With 1, 2, complete, implement.

*/
