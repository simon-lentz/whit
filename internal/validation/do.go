package validation

// Do calls the given function and recovers an IssueCollector panic. If an IssueCollector panic is recovered, its cause is returned.
// A recover of something that wasn't produced with the Error() function will result in a new panic (a rethrow).
func Do(doer func()) (collector IssueCollector) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if collector, ok = r.(IssueCollector); !ok {
				panic(r)
			}
		}
	}()
	doer()
	return
}
