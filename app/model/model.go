package model

type Service struct {
	ServiceNo        int    `json:"service_no"`
	ServiceName      string `json:"servicename"`
	Platform         string `json:"platform"`
	Env              string `json:"env"`
	Dbtype           string `json:"dbtype"`
	ServiceCategory  string `json:"servicecategory"`
	Status           string `json:"status"`
	Dba              string `json:"dba"`
	CreatedDate      string `json:"created_date"`
	UpdatedDate      string `json:"updated_date"`
	NumHost          int    `json:"num_host"`
	NumHostSkipHost  int    `json:"num_host_skip_host"`
	NumHostSkipAlert int    `json:"num_host_skip_alert"`
}

type Host struct {
	Hostname                string `json:"hostname"`
	Port                    int    `json:"port"`
	ServiceNo               int    `json:"service_no"`
	ServerGroup             string `json:"servergroup"`
	AlertskipHost           int    `json:"alertskip_host"`
	AlertskipCpuusage       int    `json:"alertskip_cpuusage"`
	AlertskipMemusage       int    `json:"alertskip_memusage"`
	AlertskipDiskspace      int    `json:"alertskip_diskspace"`
	AlertskipSysshutdown    int    `json:"alertskip_sysshutdown"`
	AlertskipSqluptime      int    `json:"alertskip_sqluptime"`
	AlertskipMirrorstate    int    `json:"alertskip_mirrorstate"`
	AlertskipMirrorrole     int    `json:"alertskip_mirrorrole"`
	AlertskipAgsyncstate    int    `json:"alertskip_agsyncstate"`
	AlertskipAgrole         int    `json:"alertskip_agrole"`
	AlertskipDboffline      int    `json:"alertskip_dboffline"`
	AlertskipDbrecov        int    `json:"alertskip_dbrecov"`
	AlertskipDbrecovpending int    `json:"alertskip_dbrecovpending"`
	AlertskipDbsusp         int    `json:"alertskip_dbsusp"`
	AlertboundCpuusage      string `json:"alertbound_cpuusage"`
	AlertboundMemusage      string `json:"alertbound_memusage"`
	AlertboundDiskspace     string `json:"alertbound_diskspace"`
	CreatedDate             string `json:"created_date"`
	UpdatedDate             string `json:"updated_date"`
}

type DBHandler interface {
	GetServiceList() []*Service
	GetServiceByServiceNo(serviceno int) []*Service
	GetHostListByServiceNo(serviceno int) []*Host
	Close()
}

func NewDBHandler() DBHandler {
	return newMySQLHandler()
}
