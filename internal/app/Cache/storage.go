package cache

const (
	WEEKLY_UPDATE = iota
	DAYLY_UPDATE
	WITHOUT_UPDATE
	WITHOUT_CACHE
)

var commandsUpdate map[string]int = map[string]int{
	"now":                WITHOUT_CACHE,
	"today":              DAYLY_UPDATE,
	"tommorow":           DAYLY_UPDATE,
	"teacher_all":        WITHOUT_UPDATE,
	"weekschedule":       WEEKLY_UPDATE,
	"weekschedule_short": WEEKLY_UPDATE,
	"week":               WEEKLY_UPDATE,
	"group":              WITHOUT_UPDATE,
	"auth":               WITHOUT_CACHE,
	"auth_teacher":       WITHOUT_CACHE,
	"deauth":             WITHOUT_CACHE,
	"deauth_teacher":     WITHOUT_CACHE,
	"subscription":       WITHOUT_CACHE,
	"subscribe":          WITHOUT_CACHE,
	"desubscribe":        WITHOUT_CACHE,
	"help":               WITHOUT_CACHE,
	"class_time":         WITHOUT_CACHE,
}
