package model

import (		
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlHandler struct {
	db *sql.DB
}

func (s *mysqlHandler) GetServiceList() []*Service {
	services := []*Service{}

	queryString := ""
	queryString += " select                                                                     "
	queryString += "     TABA.service_no                                                        "
	queryString += "   , TABA.servicename                                                       "
	queryString += "   , TABA.platform                                                          "
	queryString += "   , TABA.env                                                               "
	queryString += "   , TABA.dbtype                                                            "
	queryString += "   , TABA.servicecategory                                                   "
	queryString += "   , TABA.status                                                            "
	queryString += "   , TABA.dba                                                               "
	queryString += "   , TABA.created_date                                                      "
	queryString += "   , TABA.updated_date                                                      "
	queryString += "   , ifnull(TABB.cnt_host,0)             AS num_host                        "
	queryString += "   , ifnull(TABB.is_alertskip_host,0)    AS num_host_skip_host              "
	queryString += "   , ifnull(TABB.is_alertskip_alert,0)   AS num_host_skip_alert             "
	queryString += " from alert_service TABA                                                    "
	queryString += "       left join                                                            "
	queryString += "       (                                                                    "
	queryString += "          select                                                            "
	queryString += "              service_no                                                   "
	queryString += "            , count(hostname) as cnt_host                                   "
	queryString += "            , sum(alertskip_host) as is_alertskip_host                      "
	queryString += "            , sum(                                                          "
	queryString += "                   if (                                                     "
	queryString += "                         (                                                  "
	queryString += "                            alertskip_cpuusage                              "
	queryString += "                           +alertskip_memusage                              "
	queryString += "                           +alertskip_diskspace                             "
	queryString += "                           +alertskip_sysshutdown                           "
	queryString += "                           +alertskip_sqluptime                             "
	queryString += "                           +alertskip_mirrorstate                           "
	queryString += "                           +alertskip_mirrorrole                            "
	queryString += "                           +alertskip_agsyncstate                           "
	queryString += "                           +alertskip_agrole                                "
	queryString += "                           +alertskip_dboffline                             "
	queryString += "                           +alertskip_dbrecov                               "
	queryString += "                           +alertskip_dbrecovpending                        "
	queryString += "                           +alertskip_dbsusp                                "
	queryString += "                         ) > 0, 1, 0                                        "
	queryString += "                       )                                                    "
	queryString += "                   ) as is_alertskip_alert                                  "
	queryString += "         from alert_host                                                    "
	queryString += "         group by service_no                                               "
	queryString += "       ) TABB on TABA.service_no = TABB.service_no                        "

	//fmt.Print(queryString)
	rows, err := s.db.Query(queryString)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var service Service
		err := rows.Scan(
			&service.ServiceNo,
			&service.ServiceName,
			&service.Platform,
			&service.Env,
			&service.Dbtype,
			&service.ServiceCategory,
			&service.Status,
			&service.Dba,
			&service.CreatedDate,
			&service.UpdatedDate,
			&service.NumHost,
			&service.NumHostSkipHost,
			&service.NumHostSkipAlert,
		)
		if err != nil {
			log.Fatal(err)
		}
		services = append(services, &service)
	}

	// e, err := json.Marshal(services)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(e))

	//fmt.Println(services[0].created_date)
	return services
}

func (s *mysqlHandler) GetServiceByServiceNo(serviceno int) []*Service {
	services := []*Service{}

	queryString := ""
	queryString += " select                                                                     "
	queryString += "     TABA.service_no                                                        "
	queryString += "   , TABA.servicename                                                       "
	queryString += "   , TABA.platform                                                          "
	queryString += "   , TABA.env                                                               "
	queryString += "   , TABA.dbtype                                                            "
	queryString += "   , TABA.servicecategory                                                   "
	queryString += "   , TABA.status                                                            "
	queryString += "   , TABA.dba                                                               "
	queryString += "   , TABA.created_date                                                      "
	queryString += "   , TABA.updated_date                                                      "
	queryString += "   , ifnull(TABB.cnt_host,0)             AS num_host                        "
	queryString += "   , ifnull(TABB.is_alertskip_host,0)    AS num_host_skip_host              "
	queryString += "   , ifnull(TABB.is_alertskip_alert,0)   AS num_host_skip_alert             "
	queryString += " from alert_service TABA                                                    "
	queryString += "       left join                                                            "
	queryString += "       (                                                                    "
	queryString += "          select                                                            "
	queryString += "              service_no                                                   "
	queryString += "            , count(hostname) as cnt_host                                   "
	queryString += "            , sum(alertskip_host) as is_alertskip_host                      "
	queryString += "            , sum(                                                          "
	queryString += "                   if (                                                     "
	queryString += "                         (                                                  "
	queryString += "                            alertskip_cpuusage                              "
	queryString += "                           +alertskip_memusage                              "
	queryString += "                           +alertskip_diskspace                             "
	queryString += "                           +alertskip_sysshutdown                           "
	queryString += "                           +alertskip_sqluptime                             "
	queryString += "                           +alertskip_mirrorstate                           "
	queryString += "                           +alertskip_mirrorrole                            "
	queryString += "                           +alertskip_agsyncstate                           "
	queryString += "                           +alertskip_agrole                                "
	queryString += "                           +alertskip_dboffline                             "
	queryString += "                           +alertskip_dbrecov                               "
	queryString += "                           +alertskip_dbrecovpending                        "
	queryString += "                           +alertskip_dbsusp                                "
	queryString += "                         ) > 0, 1, 0                                        "
	queryString += "                       )                                                    "
	queryString += "                   ) as is_alertskip_alert                                  "
	queryString += "         from alert_host                                                    "
	queryString += "         where service_no = ?                                               "
	queryString += "         group by service_no                                                "
	queryString += "       ) TABB on TABA.service_no = TABB.service_no                          "
	queryString += "  where TABA.service_no = ?													"

	//fmt.Print(queryString)
	rows, err := s.db.Query(queryString, serviceno, serviceno)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var service Service
		err := rows.Scan(
			&service.ServiceNo,
			&service.ServiceName,
			&service.Platform,
			&service.Env,
			&service.Dbtype,
			&service.ServiceCategory,
			&service.Status,
			&service.Dba,
			&service.CreatedDate,
			&service.UpdatedDate,
			&service.NumHost,
			&service.NumHostSkipHost,
			&service.NumHostSkipAlert,
		)
		if err != nil {
			log.Fatal(err)
		}
		services = append(services, &service)
	}

	// e, err := json.Marshal(services)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(e))

	//fmt.Println(services[0].CreatedDate)
	return services
}

func (s *mysqlHandler) GetHostListByServiceNo(serviceno int) []*Host {
	hosts := []*Host{}

	queryString := ""
	queryString += " select                          		                          "
	queryString += "     hostname                                                 "
	queryString += "   , port                                                     "
	queryString += "   , service_no                                               "
	queryString += "   , servergroup                                              "
	queryString += "   , alertskip_host                                           "
	queryString += "   , alertskip_cpuusage                                       "
	queryString += "   , alertskip_memusage                                       "
	queryString += "   , alertskip_diskspace                                      "
	queryString += "   , alertskip_sysshutdown                                    "
	queryString += "   , alertskip_sqluptime                                      "
	queryString += "   , alertskip_mirrorstate                                    "
	queryString += "   , alertskip_mirrorrole                                     "
	queryString += "   , alertskip_agsyncstate                                    "
	queryString += "   , alertskip_agrole                                         "
	queryString += "   , alertskip_dboffline                                      "
	queryString += "   , alertskip_dbrecov                                        "
	queryString += "   , alertskip_dbrecovpending                                 "
	queryString += "   , alertskip_dbsusp                                         "
	queryString += "   , ifnull(alertbound_cpuusage,'')  as alertbound_cpuusage   "
	queryString += "   , ifnull(alertbound_memusage,'')  as alertbound_memusage   "
	queryString += "   , ifnull(alertbound_diskspace,'')  as alertbound_diskspace "
	queryString += "   , created_date                                             "
	queryString += "   , updated_date                                             "
	queryString += " from alert_host                                              "
	queryString += " where service_no = ?                                         "
	queryString += " order by servergroup, hostname                               "

	//fmt.Print(queryString)
	rows, err := s.db.Query(queryString, serviceno)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var host Host
		err := rows.Scan(
			&host.Hostname,
			&host.Port,
			&host.ServiceNo,
			&host.ServerGroup,
			&host.AlertskipHost,
			&host.AlertskipCpuusage,
			&host.AlertskipMemusage,
			&host.AlertskipDiskspace,
			&host.AlertskipSysshutdown,
			&host.AlertskipSqluptime,
			&host.AlertskipMirrorstate,
			&host.AlertskipMirrorrole,
			&host.AlertskipAgsyncstate,
			&host.AlertskipAgrole,
			&host.AlertskipDboffline,
			&host.AlertskipDbrecov,
			&host.AlertskipDbrecovpending,
			&host.AlertskipDbsusp,
			&host.AlertboundCpuusage,
			&host.AlertboundMemusage,
			&host.AlertboundDiskspace,
			&host.CreatedDate,
			&host.UpdatedDate,
		)
		if err != nil {
			log.Fatal(err)
		}

		hosts = append(hosts, &host)
	}

	// e, err := json.Marshal(hosts)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(e))

	return hosts
}

func (s *mysqlHandler) Close() {
	s.db.Close()
}

func newMySQLHandler() DBHandler {
	conn, err := sql.Open("mysql", "tommssql:tommssql2#@tcp(10.127.42.242:20306)/tommeta_mssql")
	if err != nil {
		panic(err)
	}
	return &mysqlHandler{db: conn}
}
