package isdayoff

// CountryCode type
type CountryCode string

const (
	// CountryCodeBelarus BY
	CountryCodeBelarus CountryCode = "by"
	// CountryCodeKazakhstan KZ
	CountryCodeKazakhstan CountryCode = "kz"
	// CountryCodeRussia RU
	CountryCodeRussia CountryCode = "ru"
	// CountryCodeUkraine UA
	CountryCodeUkraine CountryCode = "ua"
	// CountryCodeUSA USA
	CountryCodeUSA CountryCode = "us"
	// CountryCodeUzbekistan Uz
	CountryCodeUzbekistan CountryCode = "uz"
	// CountryCodeTurkey Turkey
	CountryCodeTurkey CountryCode = "tr"
)

// DayType type
type DayType string

// YearType type
type YearType string

// ErrorCode type
type ErrorCode string

const (
	// DayTypeWorking working day
	DayTypeWorking DayType = "0"
	// DayTypeNonWorking non working day
	DayTypeNonWorking DayType = "1"
	// DayTypeHaldHoliday half holiday
	DayTypeHaldHoliday DayType = "2"
	// DayTypeWorkingCovid working day for Covid
	DayTypeWorkingCovid DayType = "4"

	// YearTypeNotLeap leap year
	YearTypeNotLeap YearType = "0"
	// YearTypeLeap non leap year
	YearTypeLeap YearType = "1"

	// ErrorCodeWrongDate wrong date err
	ErrorCodeWrongDate ErrorCode = "100"
	// ErrorCodeNotFound not found err
	ErrorCodeNotFound ErrorCode = "101"
	// ErrorCodeInternalError internal error
	ErrorCodeInternalError ErrorCode = "199"
)
