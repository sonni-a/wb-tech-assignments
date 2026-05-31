package grep

// Options configures grep behaviour.
type Options struct {
	AfterContext  int
	BeforeContext int
	Count         bool
	IgnoreCase    bool
	Invert        bool
	FixedString   bool
	LineNumber    bool
}
