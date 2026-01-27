package model

type GridChunk struct {
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	FileID int64  `gorm:"column:file_id;index:idx_file_id_serial;comment:文件ID"`
	Serial int    `gorm:"column:serial;index:idx_file_id_serial;comment:文件分片序号"`
	Data   []byte `gorm:"column:data;comment:文件内容分片"`
}

func (GridChunk) TableName() string {
	return "gridfs_chunk"
}
