package objectsmispformat

import "sync"

// ListFormatsMISP содержит описание типов добавляемых в MISP и их порядок добавления.
// По результатам добавления Event, MISP возвращает id котрый необходим как для добавления
// следующих типов объектов MISP, так и для добавления этого значения в поле
// 'customFields' TheHive. Не все из этих объектов могут сразу добавлятся в MISP 'как есть',
// некоторые из них подлежат дополнительной обработке, см. обработчик для каждого из объектов.
// После добавления всех объектов, событие MISP необходимо опобликовать, как это сделать
// см. обработчик публикации.
type ListFormatsMISP struct {
	ID         string
	Event      *EventsMispFormat
	Reports    *EventReports
	Attributes []*AttributesMispFormat
	Objects    map[int]*ObjectsMispFormat
	ObjectTags *ListEventObjectTags
}

// EventsMispFormat формат MISP типа Events для взаимодействия с API MISP
type EventsMispFormat struct {
	OrgId              string `json:"org_id"`
	OrgcId             string `json:"orgc_id"`
	Distribution       string `json:"distribution"` //цифры в виде строки из списка
	Info               string `json:"info"`
	Uuid               string `json:"uuid"`
	Date               string `json:"date"`
	Analysis           string `json:"analysis"` //цифры в виде строки из списка
	AttributeCount     string `json:"attribute_count"`
	Timestamp          string `json:"timestamp"`
	SharingGroupId     string `json:"sharing_group_id"`
	ThreatLevelId      string `json:"threat_level_id"`    //цифры в виде строки из списка
	PublishTimestamp   string `json:"publish_timestamp"`  //по умолчанию "0"
	SightingTimestamp  string `json:"sighting_timestamp"` //по умолчанию "0"
	ExtendsUuid        string `json:"extends_uuid"`
	EventCreatorEmail  string `json:"event_creator_email"`
	Published          bool   `json:"published"`
	ProposalEmailLock  bool   `json:"proposal_email_lock"`
	Locked             bool   `json:"locked"`
	DisableCorrelation bool   `json:"disable_correlation"`
}

// EventReports формат MISP типа Reports для взаимодействия с API MISP
type EventReports struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	Distribution string `json:"distribution"`
}

type ListAttributesMispFormat struct {
	mutex      sync.Mutex
	attributes map[int]AttributesMispFormat
}

// AttributesMispFormat формат MISP типа Attributes для взаимодействия с API MISP
type AttributesMispFormat struct {
	EventId            string `json:"event_id"`
	ObjectId           string `json:"object_id"`
	ObjectRelation     string `json:"object_relation"`
	Category           string `json:"category"` //содержит одно из значений предустановленного списка
	Type               string `json:"type"`     //содержит одно из значений предустановленного списка
	Value              string `json:"value"`
	Uuid               string `json:"uuid"`
	Timestamp          string `json:"timestamp"`    //по умолчанию "0"
	Distribution       string `json:"distribution"` //цифры в виде строки из списка
	SharingGroupId     string `json:"sharing_group_id"`
	Comment            string `json:"comment"`
	FirstSeen          string `json:"first_seen"`
	LastSeen           string `json:"last_seen"`
	ToIds              bool   `json:"to_ids"`
	Deleted            bool   `json:"deleted"`
	DisableCorrelation bool   `json:"disable_correlation"`
}

// GalaxyClustersMispFormat формат MISP типа GalaxyClusters для взаимодействия с API MISP
type GalaxyClustersMispFormat struct {
	Authors        []string                  `json:"authors"`
	GalaxyElement  []GalaxyElementMispFormat `json:"GalaxyElement"`
	Id             string                    `json:"id"`
	Uuid           string                    `json:"uuid"`
	CollectionUuid string                    `json:"collection_uuid"`
	Type           string                    `json:"type"`
	Value          string                    `json:"value"`
	TagName        string                    `json:"tag_name"`
	Description    string                    `json:"description"`
	GalaxyId       string                    `json:"galaxy_id"`
	Source         string                    `json:"source"`
	Version        string                    `json:"version"`
	Distribution   string                    `json:"distribution"` //цифры в виде строки из списка
	SharingGroupId string                    `json:"sharing_group_id"`
	OrgId          string                    `json:"org_id"`
	OrgcId         string                    `json:"orgc_id"`
	ExtendsUuid    string                    `json:"extends_uuid"`
	ExtendsVersion string                    `json:"extends_version"`
	Default        bool                      `json:"default"`
	Locked         bool                      `json:"locked"`
	Published      bool                      `json:"published"`
	Deleted        bool                      `json:"deleted"`
}

// GalaxyElementMispFormat формат MISP типа Galaxy для взаимодействия с API MISP
type GalaxyElementMispFormat struct {
	Id              string `json:"id"`
	GalaxyClusterId string `json:"galaxy_cluster_id"`
	Key             string `json:"key"`
	Value           string `json:"value"`
}

// UsersMispFormat формат MISP типа Users для взаимодействия с API MISP
type UsersMispFormat struct {
	OrgId         string `json:"org_id"`
	ServerId      string `json:"server_id"`
	Email         string `json:"email"`
	Authkey       string `json:"authkey"`
	InvitedBy     string `json:"invited_by"`
	Gpgkey        string `json:"gpgkey"`
	CertifPublic  string `json:"certif_public"`
	NidsSid       string `json:"nids_sid"`
	Newsread      string `json:"newsread"`
	RoleId        string `json:"role_id"`
	ChangePw      string `json:"change_pw"`
	Expiration    string `json:"expiration"`
	CurrentLogin  string `json:"current_login"`
	LastLogin     string `json:"last_login"`
	DateCreated   string `json:"date_created"`
	DateModified  string `json:"date_modified"`
	Autoalert     bool   `json:"autoalert"`
	Termsaccepted bool   `json:"termsaccepted"`
	Contactalert  bool   `json:"contactalert"`
	Disabled      bool   `json:"disabled"`
	ForceLogout   bool   `json:"force_logout"`
}

// OrganisationsMispFormat формата MISP типа Organisations для взаимодействия с API MISP
type OrganisationsMispFormat struct {
	RestrictedToDomain []string `json:"restricted_to_domain"`
	Name               string   `json:"name"`
	DateCreated        string   `json:"date_created"`
	DateModified       string   `json:"date_modified"`
	Description        string   `json:"description"`
	Type               string   `json:"type"`
	Nationality        string   `json:"nationality"`
	Sector             string   `json:"sector"`
	CreatedBy          string   `json:"created_by"`
	Uuid               string   `json:"uuid"`
	Contacts           string   `json:"contacts"`
	Landingpage        string   `json:"landingpage"`
	UserCount          string   `json:"user_count"`
	CreatedByEmail     string   `json:"created_by_email"`
	Local              bool     `json:"local"`
}

// ServersMispFormat формат MISP типа Servers для взаимодействия с API MISP
type ServersMispFormat struct {
	Name                string `json:"name"`
	Url                 string `json:"url"`
	Authkey             string `json:"authkey"`
	OrgId               string `json:"org_id"`
	Lastpulledid        string `json:"lastpulledid"`
	Lastpushedid        string `json:"lastpushedid"`
	Organization        string `json:"organization"`
	RemoteOrgId         string `json:"remote_org_id"`
	PullRules           string `json:"pull_rules"`
	PushRules           string `json:"push_rules"`
	CertFile            string `json:"cert_file"`
	ClientCertFile      string `json:"client_cert_file"`
	Priority            string `json:"priority"`
	Push                bool   `json:"push"`
	Pull                bool   `json:"pull"`
	PushSightings       bool   `json:"push_sightings"`
	PushGalaxyClusters  bool   `json:"push_galaxy_clusters"`
	PullGalaxyClusters  bool   `json:"pull_galaxy_clusters"`
	PublishWithoutEmail bool   `json:"publish_without_email"`
	UnpublishEvent      bool   `json:"unpublish_event"`
	SelfSigned          bool   `json:"self_signed"`
	Internal            bool   `json:"internal"`
	SkipProxy           bool   `json:"skip_proxy"`
	CachingEnabled      bool   `json:"caching_enabled"`
	CacheTimestamp      bool   `json:"cache_timestamp"`
}

// FeedsMispFormat формат MISP типа Feeds для взаимодействия с API MISP
type FeedsMispFormat struct {
	Name            string `json:"name"`
	Provider        string `json:"provider"`
	Url             string `json:"url"`
	Rules           string `json:"rules"`
	Distribution    string `json:"distribution"` //цифры в виде строки из списка
	SharingGroupId  string `json:"sharing_group_id"`
	TagId           string `json:"tag_id"`
	SourceFormat    string `json:"source_format"`
	EventId         string `json:"event_id"`
	InputSource     string `json:"input_source"`
	Headers         string `json:"headers"`
	OrgcId          string `json:"orgc_id"`
	Enabled         bool   `json:"enabled"`
	FixedEvent      bool   `json:"fixed_event"`
	DeltaMerge      bool   `json:"delta_merge"`
	Publish         bool   `json:"publish"`
	OverrideIds     bool   `json:"override_ids"`
	DeleteLocalFile bool   `json:"delete_local_file"`
	LookupVisible   bool   `json:"lookup_visible"`
	CachingEnabled  bool   `json:"caching_enabled"`
	ForceToIds      bool   `json:"force_to_ids"`
}

// TagsMispFormat формат MISP типа Tags для взаимодействия с API MISP
type TagsMispFormat struct {
	Name           string `json:"name"`
	Colour         string `json:"colour"`
	OrgId          string `json:"org_id"`
	UserId         string `json:"user_id"`
	NumericalValue string `json:"numerical_value"`
	Inherited      int    `json:"inherited"`
	HideTag        bool   `json:"hide_tag"`
	IsGalaxy       bool   `json:"is_galaxy"`
	Exportable     bool   `json:"exportable"`
	IsCustomGalaxy bool   `json:"is_custom_galaxy"`
}

// EventObjectTagsMispFormat формат MISP для загрузки в event.object.tags
type EventObjectTagsMispFormat struct {
	Event string `json:"event"`
	Tag   string `json:"tag"`
}

// UsersSettingsMispFormat формат MISP типа Users для данных приходящих из API MISP
// на GET запрос типа /admin/users
type UsersSettingsMispFormat struct {
	User         UserSettingsMispFormat         `json:"User"`
	Organisation OrganisationSettingsMispFormat `json:"Organisation"`
	Role         RoleSettingsMispFormat         `json:"Role"`
}

// UserSettingsMispFormat формат сообщения типа 'User' приходящего от MISP на запрос /admin/users
// так как весь перечень информации о пользователе в настоящее время не нужен
// 'лишние' свойства отключены
type UserSettingsMispFormat struct {
	Id           string `json:"id"`
	OrgId        string `json:"org_id"`
	ServerId     string `json:"server_id"`
	Email        string `json:"email"`
	Authkey      string `json:"authkey"`
	RoleId       string `json:"role_id"`
	CurrentLogin string `json:"current_login"`
	//InvitedBy     string `json:"invited_by"`
	//Gpgkey        string `json:"gpgkey"`
	//CertifPublic  string `json:"certif_public"`
	//NidsSid       string `json:"nids_sid"`
	//Newsread      string `json:"newsread"`
	//Expiration    string `json:"expiration"`
	//LastLogin     string `json:"last_login"`
	//LastApiAccess string `json:"last_api_access"`
	//DateCreated   string `json:"date_created"`
	//DateModified  string `json:"date_modified"`
	//ChangePw      string `json:"change_pw"`
	//Autoalert     bool   `json:"autoalert"`
	//Termsaccepted bool   `json:"termsaccepted"`
	//Contactalert  bool   `json:"contactalert"`
	//Disabled      bool   `json:"disabled"`
	//ForceLogout   bool   `json:"force_logout"`
}

// RoleSettingsMispFormat формат сообщения типа 'Role' приходящего от MISP на запрос /admin/users
type RoleSettingsMispFormat struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PermAuth     bool   `json:"perm_auth"`
	PermSiteAmin bool   `json:"perm_site_admin"`
}

// OrganisationSettingsMispFormat формат сообщения типа 'Organisation' приходящего от MISP на запрос /admin/users
type OrganisationSettingsMispFormat struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// ListObjectsMispFormat формат сообщения типа 'Objects' (нет в спецификации API)
// формируется на основе содержимого поля observables получаемого от TheHive
// которое дополнительно содержит поле attachment
type ListObjectsMispFormat struct {
	mutex   sync.Mutex
	objects map[int]ObjectsMispFormat
}

type ObjectsMispFormat struct {
	Attribute       ListAttribute `json:"Attribute"`
	ID              string        `json:"id"`
	TemplateUUID    string        `json:"template_uuid"`
	TemplateVersion string        `json:"template_version"`
	FirstSeen       string        `json:"first_seen"`
	Timestamp       string        `json:"timestamp"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	EventId         string        `json:"event_id"`
	MetaCategory    string        `json:"meta-category"`
	Distribution    string        `json:"distribution"`
}

type ListAttribute []AttributeMispFormat

type AttributeMispFormat struct {
	Category       string `json:"category"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	Distribution   string `json:"distribution"`
	ObjectRelation string `json:"object_relation"`
}

// ListEventObjectTags временное хранилище для тегов полученных из event.object.tags
type ListEventObjectTags []string
