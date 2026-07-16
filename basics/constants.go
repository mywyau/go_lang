package basics

/*
Constants are values that cannot change after they are declared.

Constants can contain:
- strings
- booleans
- integers
- floating-point numbers
- rune values

Constants cannot contain values that are only known at runtime,
such as function results or user input.
*/

// Lowercase names are private to this package.
const pi = 3.14

// Uppercase names are exported and can be used by other packages.
const Role = "Admin"

// String constant.
const ApplicationName = "Go Learning Project"

// Boolean constant.
const DebugMode = true

// Integer constant.
const MaximumRetries = 3

// Floating-point constant.
const TaxRate = 0.20

// Rune constant.
// A rune represents a Unicode character.
const Grade = 'A'

// Multiple constants can be grouped together.
const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)

// Constants can use values from other constants.
const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	SecondsPerHour   = SecondsPerMinute * MinutesPerHour
)

// An untyped numeric constant can be used with different numeric types.
const DefaultScore = 100

func ConstantExamples() {
	var scoreInt int = DefaultScore
	var scoreInt64 int64 = DefaultScore
	var scoreFloat float64 = DefaultScore

	_, _, _ = scoreInt, scoreInt64, scoreFloat
}

// A typed constant is restricted to its declared type.
const MaximumUsers int = 1000

// iota creates incrementing integer constants.
// It starts at 0 inside each const block.
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// iota is useful for enum-like values.
type UserRole int

const (
	RoleGuest UserRole = iota
	RoleMember
	RoleModerator
	RoleAdministrator
)

// The blank identifier can skip an iota value.
const (
	_  = iota // skip 0
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

// Constants can be used in array lengths because their values
// are known at compile time.
const ArraySize = 5

var Numbers [ArraySize]int

/*
These are not valid constants:

const CurrentTime = time.Now()
const UserName = fmt.Scanln()
const Values = []int{1, 2, 3}

The results of functions, slices, maps, structs, and other runtime
values generally cannot be declared as constants.
*/
