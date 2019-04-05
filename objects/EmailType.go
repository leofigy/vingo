package objects

// Enum for EmailType
type EmailType int
var etypes = [...]string{
	"html", 
	"multipart", 
	"plaintext"}

const(
	html EmailType = iota + 1
	multipart
	plaintext
)

func (t EmailType) String() string {
	return etypes[t]
}

