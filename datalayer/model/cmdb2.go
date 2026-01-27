package model

// Cmdb2 资产
//
// 物理机：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=9e7fb369-3695-406d-85ae-aa51579f94aa
// 虚拟机：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=20ab7bdd-beeb-4031-bb46-ec40986592e3
// 容器：https://oa-pan.eastmoney.com/ddwiki/space/doc?spaceId=15&fileUuid=3ce158ee-31f8-4755-b7ad-0fb5c1a24ffb
type Cmdb2 struct {
	ID                      int64       `json:"-"                                    gorm:"column:id;primaryKey"`
	Inet                    string      `json:"inet"                                 gorm:"column:inet;size:20"`
	AppCluster              string      `json:"app_cluster,omitempty"                gorm:"column:app_cluster"`
	AppDuty                 Cmdb2Duties `json:"app_duty,omitempty"                   gorm:"column:app_duty;serializer:json"`
	AppID                   string      `json:"appid,omitempty"                      gorm:"column:appid"`
	AppName                 string      `json:"appname,omitempty"                    gorm:"column:appname"`
	AutoRenew               string      `json:"auto_renew,omitempty"                 gorm:"column:auto_renew"`
	AssetID                 string      `json:"asset_id,omitempty"                   gorm:"column:asset_id"`
	AssetStatus             string      `json:"asset_status,omitempty"               gorm:"column:asset_status"`
	BaoleijiIdentity        string      `json:"baoleiji_identity,omitempty"          gorm:"column:baoleiji_identity"`
	BeetleService           string      `json:"beetle_service,omitempty"             gorm:"column:beetle_service"`
	BillingType             string      `json:"billing_type,omitempty"               gorm:"column:billing_type"`
	Brand                   string      `json:"brand,omitempty"                      gorm:"column:brand"`
	Business                string      `json:"business,omitempty"                   gorm:"column:business"`
	BusinessEnv             string      `json:"business_env,omitempty"               gorm:"column:business_env"`
	ChargeMode              string      `json:"charge_mode,omitempty"                gorm:"column:charge_mode"`
	CIType                  string      `json:"ci_type,omitempty"                    gorm:"column:ci_type"`
	CmcIP                   string      `json:"cmc_ip,omitempty"                     gorm:"column:cmc_ip"`
	CncIP                   string      `json:"cnc_ip,omitempty"                     gorm:"column:cnc_ip"`
	Comment                 string      `json:"comment,omitempty"                    gorm:"column:comment"`
	CostDepCasID            string      `json:"cost_dep_cas_id,omitempty"            gorm:"column:cost_dep_cas_id"`
	CPU                     string      `json:"cpu,omitempty"                        gorm:"column:cpu"`
	CPUCount                int         `json:"cpu_count,omitempty"                  gorm:"column:cpu_count"`
	CtcIP                   string      `json:"ctc_ip,omitempty"                     gorm:"column:ctc_ip"`
	CreateDate              string      `json:"create_date,omitempty"                gorm:"column:create_date"`
	CreatedAt               string      `json:"created_at,omitempty"                 gorm:"column:created_at"`
	CreatedTime             string      `json:"created_time,omitempty"               gorm:"column:created_time"`
	Deleted                 string      `json:"deleted,omitempty"                    gorm:"column:deleted"`
	Department              string      `json:"department,omitempty"                 gorm:"column:department"`
	Description             string      `json:"description,omitempty"                gorm:"column:description"`
	DeviceSpec              string      `json:"device_spec,omitempty"                gorm:"column:device_spec"`
	DockerCPUCount          string      `json:"docker_cpu_count,omitempty"           gorm:"column:docker_cpu_count"`
	Env                     string      `json:"env,omitempty"                        gorm:"column:env"`
	ExpDate                 string      `json:"exp_date,omitempty"                   gorm:"column:exp_date"`
	ExpiredAt               string      `json:"expired_at,omitempty"                 gorm:"column:expired_at"`
	ExternalID              string      `json:"external_id,omitempty"                gorm:"column:external_id"`
	FloatIP                 []string    `json:"float_ip,omitempty"                   gorm:"column:float_ip;serializer:json"`
	HardDisk                string      `json:"harddisk,omitempty"                   gorm:"column:harddisk"`
	HostIP                  []string    `json:"host_ip,omitempty"                    gorm:"column:host_ip;serializer:json"`
	HostType                string      `json:"host_type,omitempty"                  gorm:"column:host_type"`
	Hostname                string      `json:"hostname,omitempty"                   gorm:"column:hostname"`
	HostSN                  string      `json:"host_sn,omitempty"                    gorm:"column:host_sn"`
	HyperThreading          string      `json:"hyper_threading,omitempty"            gorm:"column:hyper_threading"`
	IDC                     string      `json:"idc,omitempty"                        gorm:"column:idc"`
	IloIP                   string      `json:"ilo_ip,omitempty"                     gorm:"column:ilo_ip"`
	Image                   string      `json:"image,omitempty"                      gorm:"column:image"`
	ImageVersion            string      `json:"image_version,omitempty"              gorm:"column:image_version"`
	InstanceID              string      `json:"instance_id,omitempty"                gorm:"column:instance_id"`
	IPv6                    []string    `json:"ipv6,omitempty"                       gorm:"column:ipv6;serializer:json"`
	ImportedAt              string      `json:"imported_at,omitempty"                gorm:"column:imported_at"`
	InScalingGroup          string      `json:"in_scaling_group,omitempty"           gorm:"column:in_scaling_group"`
	InstanceType            string      `json:"instance_type,omitempty"              gorm:"column:instance_type"`
	InternetMaxBandwidthOut int         `json:"internet_max_bandwidth_out,omitempty" gorm:"column:internet_max_bandwidth_out"`
	K8sCluster              string      `json:"k8s_cluster,omitempty"                gorm:"column:k8s_cluster"`
	KernelVersion           string      `json:"kernel_version,omitempty"             gorm:"column:kernel_version"`
	LogicCPUCount           int         `json:"logic_cpu_count,omitempty"            gorm:"column:logic_cpu_count"`
	MinionNotCheck          string      `json:"minion_not_check,omitempty"           gorm:"column:minion_not_check"`
	Name                    string      `json:"name,omitempty"                       gorm:"column:name"`
	Namespace               string      `json:"namespace,omitempty"                  gorm:"column:namespace"`
	NetOpen                 string      `json:"net_open,omitempty"                   gorm:"column:net_open"`
	OpDuty                  Cmdb2OpDuty `json:"op_duty,omitempty"                    gorm:"column:op_duty;serializer:json"`
	OpDutyBackup            Cmdb2Duties `json:"op_duty.backup,omitempty"             gorm:"column:op_duty_backup;serializer:json"`
	OpDutyMain              Cmdb2Duties `json:"op_duty.main,omitempty"               gorm:"column:op_duty_main;serializer:json"`
	OpDutyStandby           Cmdb2Duties `json:"op_duty.standby,omitempty"            gorm:"column:op_duty_standby;serializer:json"`
	OsArch                  string      `json:"os_arch,omitempty"                    gorm:"column:os_arch"`
	OsType                  string      `json:"os_type,omitempty"                    gorm:"column:os_type"`
	OsVersion               string      `json:"os_version,omitempty"                 gorm:"column:os_version"`
	PowerStates             string      `json:"power_states,omitempty"               gorm:"column:power_states"`
	PrivateIP               []string    `json:"private_ip,omitempty"                 gorm:"column:private_ip;serializer:json"`
	PublicCloudID           string      `json:"id,omitempty"                         gorm:"column:public_cloud_id"`
	PublicCloudIDC          string      `json:"public_cloud_idc,omitempty"           gorm:"column:public_cloud_idc"`
	PrivateCloudIP          string      `json:"private_cloud_ip,omitempty"           gorm:"column:private_cloud_ip"`
	PrivateCloudType        string      `json:"private_cloud_type,omitempty"         gorm:"column:private_cloud_type"`
	Rack                    string      `json:"rack,omitempty"                       gorm:"column:rack"`
	RAID                    string      `json:"raid,omitempty"                       gorm:"column:raid"`
	RAM                     string      `json:"ram,omitempty"                        gorm:"column:ram"`
	RAMSize                 string      `json:"ram_size,omitempty"                   gorm:"column:ram_size"`
	RdDutyMain              Cmdb2Duties `json:"rd_duty.main,omitempty"               gorm:"column:rd_duty_main;serializer:json"`
	RdDutyMember            Cmdb2Duties `json:"rd_duty.member,omitempty"             gorm:"column:rd_duty_member;serializer:json"`
	Region                  string      `json:"region,omitempty"                     gorm:"column:region"`
	ResourceLimits          string      `json:"resource_limits,omitempty"            gorm:"column:resource_limits"`
	ResourceRequests        string      `json:"resource_requests,omitempty"          gorm:"column:resource_requests"`
	SecurityInfo            string      `json:"security_info,omitempty"              gorm:"column:security_info"`
	Server                  string      `json:"server,omitempty"                     gorm:"column:server"`
	ServerRoom              string      `json:"server_room,omitempty"                gorm:"column:server_room"`
	SN                      string      `json:"sn,omitempty"                         gorm:"column:sn"`
	ShutdownBehavior        string      `json:"shutdown_behavior,omitempty"          gorm:"column:shutdown_behavior"`
	ShutdownMode            string      `json:"shutdown_mode,omitempty"              gorm:"column:shutdown_mode"`
	Status                  string      `json:"status,omitempty"                     gorm:"column:status"`
	SysDuty                 Cmdb2Duties `json:"sys_duty,omitempty"                   gorm:"column:sys_duty;serializer:json"`
	Tags                    []string    `json:"tags,omitempty"                       gorm:"column:tags;serializer:json"`
	Throughput              int         `json:"throughput,omitempty"                 gorm:"column:throughput" `
	TradeType               string      `json:"trade_type,omitempty"                 gorm:"column:trade_type" `
	UpdateTime              string      `json:"update_time,omitempty"                gorm:"column:update_time"`
	UpdatedAt               string      `json:"updated_at,omitempty"                 gorm:"column:updated_at" `
	Use                     string      `json:"use,omitempty"                        gorm:"column:use"`
	UUID                    string      `json:"uuid,omitempty"                       gorm:"column:uuid"`
	VCPUCount               int         `json:"vcpu_count,omitempty"                 gorm:"column:vcpu_count"`
	VMemSize                int         `json:"vmem_size,omitempty"                  gorm:"column:vmem_size"`
	VserverType             string      `json:"vserver_type,omitempty"               gorm:"column:vserver_type"`
	ZabbixNotCheck          string      `json:"zabbix_not_check,omitempty"           gorm:"column:zabbix_not_check"`
	Zone                    string      `json:"zone,omitempty"                       gorm:"column:zone"`

	// 下面的字段定义和 cmdb2 接口文档描述不一致，或者是会返回多种类型。
	// 这些字段无关紧要，目前为了不影响系统间交互，就不解析存储这些字段。
	// cmdb 系统部分字段定义比较混乱，同一个字段，返回的类型都可能不一样。
	// AppLoadType      []string           `bson:"app_load_type,omitempty"      json:"app_load_type,omitempty"`
	// SecurityRisk     int                `bson:"security_risk,omitempty"      json:"security_risk,omitempty"`
}

type Cmdb2Duty struct {
	EmployeeID string `bson:"employee_id,omitempty" json:"employee_id,omitempty"`
	Nickname   string `bson:"nickname,omitempty"    json:"nickname,omitempty"`
	UID        int    `bson:"uid,omitempty"         json:"uid,omitempty"`
	Username   string `bson:"username,omitempty"    json:"username,omitempty"`
	UType      string `bson:"utype,omitempty"       json:"utype,omitempty"`
}

type Cmdb2Duties []*Cmdb2Duty

type Cmdb2OpDuty struct {
	Backup  Cmdb2Duties `bson:"backup,omitempty"  json:"backup,omitempty"`
	Main    Cmdb2Duties `bson:"main,omitempty"    json:"main,omitempty"`
	Standby Cmdb2Duties `bson:"standby,omitempty" json:"standby,omitempty"`
}
