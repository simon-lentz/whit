# IDENTIFIERS
Exploration of identifiers in our knowledge data.

Excellent paper about identifiers (http://ceur-ws.org/Vol-1304/STIDS2014_T03_Emmons.pdf)

W3C:
> Resource names on the Semantic Web should fulfill two requirements: First, a description of the identified resource should be retrievable with standard Web technologies. Second, a naming scheme should not confuse things and the documents representing them.
>
>We have described two approaches that fulfill these requirements, both based on the HTTP URI scheme and protocol. One is to use the 303 HTTP status code to redirect from the resource identifier to the describing document. One is to use “hash URIs” to identify resources, exploiting the fact that hash URIs are retrieved by dropping the part after the hash and retrieving the other part.
>
>The requirement to distinguish between resources and their descriptions increases the need for coordination between multiple URIs. Some useful techniques are: embedding links to RDF data in HTML documents, using RDF statements to describe the relationship between the URIs, and using content negotiation to redirect to an appropriate description of a resource.


# UUID
We could use UUID (GUID) as identifiers. The pros are that they are (virtually) globally unique and can be independently created. The cons are that they are impossibly long and must be copy/pasted.

We could generate them by triggering an action on save in VSC that replaces all occurences of $$n with a UUID. The n part serves as a local identifier that is not saved in the db. For example:

```yaml
Persons:
    -   name: Henrik
        id: $$1
        MOTHER:
            id: $$2
    -   name: Anna Greta
        id: $$2
```
Replacement changes all $$n to the same generated UUID, thus making it easy to link within the document.

To make a link to something elsewhere the UUID would be entered. For example like this (using a short UUID, a full UUID has 32 characters (for example: f81d4fae-7dec-11d0-a765-00a0c91e6bf6).
```yaml
Persons:
    -   name: Henrik
        id: "BE6sol"
        MOTHER:
            id: "SRaJC5"
    -   name: Anna Greta
        id: "SRaJC5"
```

It is probably best to retain the local references in the generated ids in case there was an error in the input (linking up the wrong things for example). So, the generated result could be:

```yaml
Persons:
    -   name: Henrik
        id: "BE6sol.2"
        MOTHER:
            id: "SRaJC5.1"
    -   name: Anna Greta
        id: "SRaJC5.1"
```
We would strip the ".n" part when storing and references with and without ".n" would resolve to the same thing. It may not be of any practical use since an author would copy/search for a UUID anyway...

Note: The implementation in whit retains the prefix `$$` and keeps the local id as a `suffix`. This makes it easier to
do manual editing of a document.

## Hierarchical

RDF is based on URIs (or IRIs which allow unicode characters). It is best if the identifier is resolvable (possible to retrieve).
For example:
```
http://org/dept/project/class/item
tag:org,2014-10-01:dept/project/class/item
urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6
```
Where the `urn:uuid` is not resolvable since it does not define where to find the item.

## SHA1 Hashes, like Git
Git uses SHA-1 hashes as identifiers. They are 20 bytes long rendered as 40 hex digits. When referencing, it is almost always possible
to use the last 6 hex digits instead of the full 40 digit id.

We could use SHA-1 as hashes. We would compute the hash of the information for a struct and set that as its ID. While this is good for concistency tests, it probably just makes it more complicated as every change to a property would result in a new id of that object.
OTOH, if the reference was made with a query for "Exxon", "Main Street 2", and Exxon address was updated to "Main Street 1" then
the query would fail just the same way as if it had a hash of the original content as reference.

We could compute the hash on first save and then not update it if the document is modified. It would then be possible to detect if the data has changed from the original. How this would be used in practice and if it provides any value is questionable.

The urn scheme for sha-1 is `urn:sha1:`

## Go
Go modules and packages are based on a similar concept. A source file in go imports a package via a reference to the module and a path within the module. For example in `whit` you can find this in a test:
```
import (
	"testing"

	"github.com/hlindberg/testutils"
	"github.com/wyrth-io/whit/internal/utils"
	"github.com/wyrth-io/whit/internal/validation"
)
```
Then, in this source file it is possible to refer to the package that those paths refer to - for example `testutils.NewTester()`. Tools provided with go manages the dependencies, resolves them by fetching/downloading them etc. at a particular version (a "tag", a "commit" or "latest").

If all of our data is in `github.com/wyrth-io/rdata` then we could use that as our IRI. For RDF we could define our ontology such that
https://github.com/wyrth-io/rdata/ont/ElectricalService would serve up the semantics of what that is.
We can generate such an ontology from our schema.

# Scraping/downloading and Identities
If data does not come with identities (primary keys, or UUID or similar) we need to generate them.