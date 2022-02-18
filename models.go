package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

//customer has many uploads, many category, many restreamer, has one credential, one encoding profiles
type Customer struct {
	ID               int              `json:"id" gorm:"primaryKey"`
	Name             string           `json:"name"`
	EmailAddress     string           `json:"email_address"`
	Priority         int              `json:"priority"`
	KeycloakID       string           `json:"keycloak_id"`
	SplitTopic       int              `json:"split_topic"`
	SchedulerUser    string           `json:"scheduler_user"`
	JobTemplate      string           `json:"job_template"`
	InputType        int              `json:"input_type"`
	APIKeyID         string           `json:"api_key_id"`
	WatermarkPath    string           `json:"watermark_path"`
	CompressEndpoint string           `json:"compress_endpoint"`
	ClientID         string           `json:"client_id"`
	Credential       *Credential      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Uploads          []Upload         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EncodingProfile  *EncodingProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Restreamers      []Restreamer     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName overrides the table name used by Restreamer to `restreamer` from plural
func (Customer) TableName() string {
	return "customer"
}

// DefaultValues will create sane default values for a brand-new customer to be added to the db
func (c *Customer) DefaultValues(name string) {
	c.Name = fmt.Sprintf("%s_apikey", name)
	c.Priority = 1
	c.EmailAddress = fmt.Sprintf("%s@tngrm.io", name)
	c.KeycloakID = ""
	c.SplitTopic = 1
	c.SchedulerUser = "tangram"
	c.JobTemplate = "/mnt/beegfs/scripts/jobs/freemium5/split_template.aurora"
	c.InputType = 2
	c.APIKeyID = fmt.Sprintf("service-account-%s_client@placeholder.org", name)
	c.WatermarkPath = fmt.Sprintf("/mnt/nfs/scripts/jobs/watermarks/%s", name)
	c.CompressEndpoint = fmt.Sprintf("https://%s.tngrm.io", name)
	c.ClientID = fmt.Sprintf("%s_client", name)
}

type Credential struct {
	ID                       int    `json:"id" gorm:"primaryKey"`
	CustomerID               int    `json:"customer_id"`
	S3Host                   string `json:"s3_host"`
	S3Bucket                 string `json:"s3_bucket"`
	S3AccessKey              string `json:"s3_access_key"`
	S3SecretKey              string `json:"s3_secret_key"`
	SftpUsername             string `json:"sftp_username"`
	SftpPassword             string `json:"sftp_password"`
	SftpRemoteHost           string `json:"sftp_remote_host"`
	SftpRemotePort           int    `json:"sftp_remote_port"`
	SftpBasePath             string `json:"sftp_base_path"`
	Protocol                 string `json:"protocol"`
	EndpointVisualization    string `json:"endpoint_visualization"`
	NfsMount                 string `json:"nfs_mount"`
	VodesBaseURL             string `json:"vodes_base_url"`
	LiveBaseURL              string `json:"live_base_url"`
	LiveBaseURLDedicated     string `json:"live_base_url_dedicated"`
	ExternalClusters         string `json:"external_clusters"`
	RestreamerHlsBasePath    string `json:"restreamer_hls_base_path"`
	RestreamerHlstmpBasePath string `json:"restreamer_hlstmp_base_path"`
	ClientID                 string `json:"client_id"`
}

func (Credential) TableName() string {
	return "credentials"
}

// DefaultValues will produce sane default values for a brand-new credential record
func (c *Credential) DefaultValues(customerId int, customerName string) {
	c.CustomerID = customerId
	c.S3Host = ""
	c.S3Bucket = ""
	c.S3AccessKey = ""
	c.S3SecretKey = ""
	c.SftpUsername = "tangram"
	c.SftpPassword = "tangram"
	c.SftpRemoteHost = "10.15.20.90"
	c.SftpRemotePort = 22
	c.SftpBasePath = fmt.Sprintf("/mnt/tangram/%s/VMFS1/FILES/public/videos/%s/", customerName, customerName)
	c.Protocol = "sftp"
	c.EndpointVisualization = fmt.Sprintf("https://%s-l3-vod.secure.footprint.net/%s/VMFS1/FILES/public/videos/%s/", customerName, customerName, customerName)
	c.NfsMount = fmt.Sprintf("/mnt/tangram/%s/VMFS1/FILES/public/videos/%s", customerName, customerName)
	c.VodesBaseURL = fmt.Sprintf("https://%s-l3-vod.secure.footprint.net/%s", customerName, customerName)
	c.LiveBaseURL = fmt.Sprintf("https://%s-l3-live1.secure.footprint.net/%s", customerName, customerName)
	c.LiveBaseURLDedicated = fmt.Sprintf("https://%s-l3-live1.secure.footprint.net/%s", customerName, customerName)
	c.ExternalClusters = ""
	c.RestreamerHlsBasePath = fmt.Sprintf("/mnt/tangram/%s/restreamer/%s", customerName, fmt.Sprintf("%s_client", customerName))
	c.RestreamerHlstmpBasePath = fmt.Sprintf("/hlstmp/%s", fmt.Sprintf("%s_client", customerName))
	c.ClientID = fmt.Sprintf("%s_client", customerName)
}

type CustomerApiKey struct {
	ID         string `json:"id" gorm:"primaryKey"` // eg: footters_id
	APIKey     string `json:"api_key"`
	CustomerID int    `json:"customer_id"`
}

// TableName overrides the table name used by Restreamer to `restreamer` from plural
func (CustomerApiKey) TableName() string {
	return "api_keys"
}

type EncodingProfile struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Customer      string `json:"customer"`
	Hls_high_240  int8   `json:"hls_high_240p"`
	Hls_high_360  int8   `json:"hls_high_360p"`
	Hls_high_405  int8   `json:"hls_high_405p"`
	Hls_high_480  int8   `json:"hls_high_480p"`
	Hls_high_540  int8   `json:"hls_high_580p"`
	Hls_high_720  int8   `json:"hls_high_720p"`
	Hls_high_1080 int8   `json:"hls_high_1080p"`
	Hls_240       int8   `json:"hls_240p"`
	Hls_360       int8   `json:"hls_360p"`
	Hls_405       int8   `json:"hls_405p"`
	Hls_480       int8   `json:"hls_480p"`
	Hls_540       int8   `json:"hls_540p"`
	Hls_720       int8   `json:"hls_720p"`
	Hls_1080      int8   `json:"hls_1080p"`
	Mp4_high_240  int8   `json:"mp4_high_240p"`
	Mp4_high_360  int8   `json:"mp4_high_360p"`
	Mp4_high_405  int8   `json:"mp4_high_405p"`
	Mp4_high_480  int8   `json:"mp4_high_480p"`
	Mp4_high_540  int8   `json:"mp4_high_540p"`
	Mp4_high_720  int8   `json:"mp4_high_720p"`
	Mp4_high_1080 int8   `json:"mp4_high_1080p"`
	Mp4_240       int8   `json:"mp4_240p"`
	Mp4_360       int8   `json:"mp4_360p"`
	Mp4_405       int8   `json:"mp4_405p"`
	Mp4_480       int8   `json:"mp4_480p"`
	Mp4_540       int8   `json:"mp4_540p"`
	Mp4_720       int8   `json:"mp4_720p"`
	Mp4_1080      int8   `json:"mp4_1080p"`
}

func (p *EncodingProfile) DefaultValues(customerName string) {
	p.Customer = customerName
	p.Hls_high_240 = 0
	p.Hls_high_360 = 1
	p.Hls_high_405 = 0
	p.Hls_high_480 = 1
	p.Hls_high_540 = 0
	p.Hls_high_720 = 1
	p.Hls_high_1080 = 1
	p.Hls_240 = 0
	p.Hls_360 = 0
	p.Hls_405 = 0
	p.Hls_480 = 0
	p.Hls_480 = 0
	p.Hls_540 = 0
	p.Hls_720 = 0
	p.Hls_1080 = 0
	p.Mp4_high_240 = 0
	p.Mp4_high_360 = 0
	p.Mp4_high_405 = 0
	p.Mp4_high_480 = 0
	p.Mp4_high_540 = 0
	p.Mp4_high_540 = 0
	p.Mp4_high_720 = 0
	p.Mp4_high_1080 = 0
	p.Mp4_240 = 0
	p.Mp4_360 = 0
	p.Mp4_405 = 0
	p.Mp4_480 = 0
	p.Mp4_540 = 0
	p.Mp4_720 = 0
	p.Mp4_1080 = 0
}

//category has many upload
type Category struct {
	ID          int      `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Created     string   `json:"created"`
	VideoAmount int      `json:"video_amount"`
	Customer    string   `json:"customer"`
	Uploads     []Upload `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

//upload has one thumbnails, manifesturl and many tags
type Upload struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title"`
	Created       time.Time      `json:"created"`
	Email         string         `json:"email"`
	Location      string         `json:"location"`
	IpAddress     string         `json:"ip_address"`
	Size          int            `json:"size"`
	Length        int            `json:"length"`
	NginxPath     string         `json:"nginx_path"`
	FileName      string         `json:"file_name"`
	EmailSent     int            `json:"email_sent"`
	JobId         int            `json:"job_id"`
	ReceivedToken int            `json:"received_token"`
	DownloadUrl   string         `json:"download_url"` //TODO to deprecate with new models
	RemoteUrl     string         `json:"remote_url"`   //TODO deprected
	OriginalUrl   string         `json:"original_url"` //TODO to deprecate with new models
	WatermarkUrl  string         `json:"watermark_url"`
	Completed     int            `json:"completed"`
	UploadDone    int            `json:"upload_done"`
	Profiles      string         `json:"profiles"`
	Token         string         `json:"token"`
	DownloadLink  string         `json:"download_link"`
	JobFailed     int            `json:"job_failed"`
	Retries       int            `json:"retries"`
	OutputPath    string         `json:"output_path"`
	OutputType    string         `json:"output_type"`
	Notification  int            `json:"notification"`
	CustomerID    int            `json:"customer_id"`
	Thumbnail     string         `json:"thumbnail"` //TODO to deprecate with new models
	ReportedID    int            `json:"reported_id"`
	CategoryID    int            `json:"category_id"`
	Thumbnails    string         `json:"thumbnails"` //TODO to deprecate with new models
	ExtendedData  string         `json:"extended_data"`
	ClientID      string         `json:"client_id"`
	ThumbnailsURL *ThumbnailsURL `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ManifestURL   *ManifestURL   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tags          []Tag          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// TableName overrides the table name used by Upload to `upload` from plural
func (Upload) TableName() string {
	return "upload"
}

// DefaultValues will produce sane default values for a brand-new Upload record
func (v *Upload) DefaultValues(title string, location string, reporterID int, categoryID int, clientId string, filename string, Size int, jobId int) {
	v.Title = title
	v.Created = time.Now().UTC()
	v.Email = ""
	v.Location = location
	v.IpAddress = ""
	v.Size = Size
	v.Length = 0
	v.NginxPath = ""
	v.FileName = "play.mp4"
	v.EmailSent = 0
	v.JobId = jobId
	v.ReceivedToken = 1
	v.DownloadUrl = ""
	v.RemoteUrl = ""
	v.OriginalUrl = ""
	v.WatermarkUrl = ""
	v.Completed = 0
	v.UploadDone = 1
	v.Profiles = ""
	v.DownloadLink = ""
	v.JobFailed = 0
	v.Retries = 0
	v.OutputPath = ""
	v.OutputType = ""
	v.Notification = 0
	v.CustomerID = 0
	v.Thumbnail = ""
	v.ReportedID = reporterID
	v.CategoryID = categoryID
	v.Thumbnails = ""
	v.ExtendedData = "[]"
	v.ClientID = clientId
	v.Token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJBeFhhdjd6SUZ5R2Zsc2pZdTYwaEJRVFh5eGFoSkhFTm52cTJqTUZubG9zIn0.eyJqdGkiOiJjYzBlMmMzNy04NDdlLTRlNjMtYjNjOC0wYjFmZTI1NjNlM2QiLCJleHAiOjE2NDM2NzMxOTksIm5iZiI6MCwiaWF0IjoxNjQzNjM3MTk5LCJpc3MiOiJodHRwczovL2F1dGgudG5ncm0uaW8vYXV0aC9yZWFsbXMvdGFuZ3JhbSIsInN1YiI6IjU2ZDI3NjQ5LTcwNjAtNGUwZi1hNGQ1LWU1MWE4NTJkYWZlNyIsInR5cCI6IkJlYXJlciIsImF6cCI6ImZlZGVybW90b19jbGllbnQiLCJhdXRoX3RpbWUiOjAsInNlc3Npb25fc3RhdGUiOiJhZjExZjVmZi1kOThhLTRmN2UtYjA5OC1kNmViZWEwY2Y5ODciLCJhY3IiOiIxIiwic2NvcGUiOiJwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRJZCI6ImZlZGVybW90b19jbGllbnQiLCJjbGllbnRIb3N0IjoiMTAuMC4xLjEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJzZXJ2aWNlLWFjY291bnQtZmVkZXJtb3RvX2NsaWVudCIsImNsaWVudEFkZHJlc3MiOiIxMC4wLjEuMSJ9.U5W4JugQ4krGhvRl4BXVwrRePCCvuY_45Rv6IcvlJRjNfR1SE0sLjxmHxJTWOzIHsPdlEw4xFhxJOjkoJG090tTt8lb0L8dIjBIkrm7DFZRHV_-mpbYjnM5zdPqICQmJImzNSDm0f9BVGkMRtolTvI7P0WvbW3PTnQCBP4r630CYzY9sRElxgismShZWkxd00Tmx5w5JR53yvk4WpIg0kjgOKmvmIaLC4Xee9PrYr0l9-mj-ZTqhyxngJJmMiUDNgtXwJUdZx8dJHPQ-UuZtA9LpCI0rR5xTaPadHhp_GL2znVCWwVRhBd4C7ge6wWyhYcgVGJxQeHVtNFvTsk6QSA"

	//TODO: add video duration

	if filename != "" {
		v.FileName = filename
	}
}

type ThumbnailsURL struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	High1080 []string `json:"high_1080"`
	High360  []string `json:"high_360"`
	High720  []string `json:"high_720"`
	High480  []string `json:"high_480"`
	Original []string `json:"original"`
}

type ManifestURL struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	High720  string `json:"high_720"`
	High1080 string `json:"high_1080"`
	High360  string `json:"high_360"`
	High480  string `json:"high_480"`
}

type Tag struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Customer string `json:"customer"`
}

// // gpu has many restreamers

// type Gpu struct {
// 	Restreamers []Restreamer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// //gpuServer has many gpus
// type GpuServer struct {
// 	Gpus []Gpu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

//restreamer has many push history, many startupHistory, one srt, one settings
type Restreamer struct {
	Id                   int                        `json:"id" gorm:"primaryKey"`
	Enabled              int                        `json:"enabled"`
	Name                 string                     `json:"name"`
	CustomerId           int                        `json:"customer_id"`
	Dnsname              string                     `json:"dnsname"`
	MarathonUrl          string                     `json:"marathon_url"`
	NpmUrl               string                     `json:"npm_url"`
	ExternalRtmp         string                     `json:"external_rtmp"`
	LoopbackRtmp         string                     `json:"loopback_rtmp"`
	Md5Generator         *string                    `json:"md5_generator" gorm:"column:md5-generator"`
	MarathonPath         string                     `json:"marathon_path"`
	Token                string                     `json:"token"`
	DockerTemplate       string                     `json:"docker_template"`
	Status               int                        `json:"status"`
	Owner                string                     `json:"owner"`
	Title                string                     `json:"title"`
	Description          string                     `json:"description"`
	OwnerId              int                        `json:"owner_id"`
	CdnPath              string                     `json:"cdn_path"`
	ExtendedData         string                     `json:"extended_data"`
	LatestEvent          string                     `json:"latest_event"`
	Label                string                     `json:"label"`
	Dedicated            string                     `json:"dedicated"`
	StaticIp             string                     `json:"static_ip"`
	NvidiaVisibleDevices string                     `json:"nvidia_visible_devices"`
	Transcode            int                        `json:"transcode"`
	UpdateAt             time.Time                  `json:"update_at"`
	PushHistory          []RestreamerPushHistory    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartupHistory       []RestreamerStartupHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RestreamerSrt        *RestreamerSrt             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RestreamerSettings   *RestreamerSettings        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RestreamerEvent      []RestreamerEvent          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Restreamer to `restreamer` from plural
func (Restreamer) TableName() string {
	return "restreamer"
}

// DefaultValues set sane default values for a brand new restreamer
//
// the default owner should be "demo" customer
func (r *Restreamer) DefaultValues(name string, owner string, ownerId int, cdnBaseUrl string, guidEvent string, dedicatedServer string, staticIp string, nvidiaDevice string) {
	r.Enabled = 1
	r.Name = name
	r.Dnsname = fmt.Sprintf("%s.rtmp.marathon.mesos", name)
	r.MarathonUrl = "http://marathon.mesos:8080"
	r.NpmUrl = fmt.Sprintf("https://stream3.tngrm.io/%s/", name)
	r.LoopbackRtmp = "rtmp://localhost/live/live"
	r.Md5Generator = nil
	r.MarathonPath = "/rtmp/"
	//md5 generate default token
	h := uuid.New().String()
	r.Token = strings.ToUpper(h)
	//
	r.DockerTemplate = "rs_v2"
	r.ExternalRtmp = fmt.Sprintf("rtmp://ms.tngrm.io/live/live?token=%s&vhost=%s", r.Token, r.Name)
	r.Status = 3
	r.Owner = owner
	r.Title = ""
	r.Description = ""
	r.OwnerId = 0
	r.CustomerId = ownerId
	r.CdnPath = fmt.Sprintf("%s/restreamer/%s/%s/restreamer/rtmp/hls/%s/manifest.m3u8", cdnBaseUrl, fmt.Sprintf("%s_client", owner), name, guidEvent)
	r.ExtendedData = ""
	r.LatestEvent = guidEvent
	r.Label = name
	r.Dedicated = dedicatedServer
	r.StaticIp = staticIp
	r.NvidiaVisibleDevices = nvidiaDevice
	r.Transcode = 0
	r.UpdateAt = time.Now()
}

type RestreamerEvent struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	EventId      string    `json:"event_id"`
	StartAt      time.Time `json:"start_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;"`
	Active       int       `json:"active"`
	Completed    int       `json:"completed"`
	InstanceName string    `json:"instance_name"`
	CustomerId   int       `json:"customer_id"`
}

func (RestreamerEvent) TableName() string {
	return "restreamer_events"
}

func (e *RestreamerEvent) DefaultValues(customerId int, guidEvent string, eventName string, instanceName string) {
	e.InstanceName = instanceName
	e.Active = 1
	e.Completed = 0
	e.CustomerId = customerId
	e.StartAt = time.Now()
	e.UpdatedAt = time.Now()
	e.EventId = guidEvent
}

type RestreamerPushHistory struct {
	Id           int       `json:"id"`
	Label        string    `json:"label"`
	Instance     int       `json:"instance"`
	PushTo       string    `json:"push_to"`
	Protocol     string    `json:"protocol"`
	AudioChannel int       `json:"audio_channel"`
	Customer     int       `json:"customer"`
	PushDate     time.Time `json:"push_date"`
}

func (RestreamerPushHistory) TableName() string {
	return "restreamer_push_history"
}

type RestreamerSettings struct {
	Id             int    `json:"id"`
	EnableEncoding int    `json:"enable_encoding"`
	EncodingFormat string `json:"encoding_format"`
	Encoding1080   int    `json:"1080_encoding" gorm:"column:1080_encoding"`
	EnableLive24   int    `json:"enable_live24"`
	ResetTime      int    `json:"reset_time"`
}

// TableName overrides the table name used by Restreamer to `restreamer` from plural
func (RestreamerSettings) TableName() string {
	return "restreamer_settings"
}

type RestreamerSrt struct {
	Id              int     `json:"id"`
	RestreamerId    int     `json:"restreamer_id"`
	Port            int     `json:"port"`
	ExternalSrt     string  `json:"external_srt"`
	Token           string  `json:"token"`
	CdnPath         string  `json:"cdn_path"`
	Owner           string  `json:"owner"`
	LatestEvent     string  `json:"latest_event"`
	InternalIp      *string `json:"internal_ip"`
	Mode            string  `json:"mode"`
	TranscodingType *string `json:"transcoding_type"`
	StreamId        string  `json:"streamid" gorm:"column:streamid"`
	Enabled         int     `json:"enabled"`
}

// TableName overrides the table name used by Restreamer to `restreamer` from plural
func (RestreamerSrt) TableName() string {
	return "restreamer_srt"
}

func (r *RestreamerSrt) DefaultValues(newRestreamer Restreamer, lastSrtPort int, cdnBaseUrl string, owner string, instanceName string, guidEvent string) {
	r.RestreamerId = newRestreamer.Id
	r.Port = lastSrtPort + 1
	r.Token = "1235"
	r.CdnPath = fmt.Sprintf("%s/restreamer/%s/%s/restreamer/srt/hls/%s/manifest.m3u8", cdnBaseUrl, fmt.Sprintf("%s_client", owner), instanceName, guidEvent)
	r.Owner = newRestreamer.Owner
	r.LatestEvent = newRestreamer.LatestEvent
	r.InternalIp = nil
	r.Mode = "caller"
	r.TranscodingType = nil
	r.StreamId = "live"
	r.Enabled = 0
	r.ExternalSrt = "srt://srt-stream3.tngrm.io"
}

type RestreamerStartupHistory struct {
	Id           int       `json:"id"`
	RestreamerId int       `json:"restreamer_id"`
	Type         string    `json:"type"`
	Time         time.Time `json:"time"`
}
