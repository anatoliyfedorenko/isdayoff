package isdayoff

// CountryCode Коды стран
type CountryCode string

const (
	// CountryCodeBY Белорусь
	CountryCodeBY CountryCode = "by"
	// CountryCodeKZ Казахстан
	CountryCodeKZ CountryCode = "kz"
	// CountryCodeRU Россия
	CountryCodeRU CountryCode = "ru"
	// CountryCodeUA Украина
	CountryCodeUA CountryCode = "ua"
)

// DayType тип дня
type DayType string

// YearType тип года
type YearType string

// ErrorCode код ошибки
type ErrorCode string

const (
	// DayTypeWorking Рабочий день
	DayTypeWorking DayType = "0"
	// DayTypeNonWorking Нерабочий день
	DayTypeNonWorking DayType = "1"
	// DayTypeHaldHoliday Сокращённый рабочий день
	DayTypeHaldHoliday DayType = "2"
	// DayTypeWorkingCovid Рабочий день (Covid)
	DayTypeWorkingCovid DayType = "4"

	// YearTypeNotLeap Невисокосный год
	YearTypeNotLeap YearType = "0"
	// YearTypeLeap Високосный год
	YearTypeLeap YearType = "1"

	// ErrorCodeWrongDate Ошибка в дате
	ErrorCodeWrongDate ErrorCode = "100"
	// ErrorCodeNotFound Данные не найдены
	ErrorCodeNotFound ErrorCode = "101"
	// ErrorCodeInternalError Ошибка сервиса
	ErrorCodeInternalError ErrorCode = "199"
)
