package model

type MinionSubstanceExclude struct {
	ID          int64 `json:"id,string"           gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID    int64 `json:"minion_id,string"    gorm:"column:minion_id;uniqueIndex:uk_minion_id_substance_id;"`
	SubstanceID int64 `json:"substance_id,string" gorm:"column:substance_id;uniqueIndex:uk_minion_id_substance_id;"`
}

func (MinionSubstanceExclude) TableName() string {
	return "minion_substance_exclude"
}

type MinionSubstanceExcludes []*MinionSubstanceExclude

func (mss MinionSubstanceExcludes) Filter(subs Substances) Substances {
	if len(subs) == 0 || len(mss) == 0 {
		return subs
	}

	result := make(Substances, 0, 10)
	index := make(map[int64]struct{}, len(mss))
	for _, m := range mss {
		index[m.SubstanceID] = struct{}{}
	}
	for _, e := range subs {
		if _, exists := index[e.ID]; exists {
			continue
		}
		result = append(result, e)
	}

	return result
}
