package model

import "strings"

type MinionTag struct {
	Name string        `bson:"name" json:"name"`
	Type MinionTagType `bson:"type" json:"type"`
}

type MinionTags []*MinionTag

func (mts MinionTags) Deduplicate() MinionTags {
	return nil
}

// DeleteByManual 人工操作删除标签。
// 人工操作下，除系统标签不可删除，其它类型都可删除。
func (mts MinionTags) DeleteByManual(tags ...string) MinionTags {
	return nil
}

// DeleteByMinion agent 删除标签。
// agent 操作下，仅可删除
func (mts MinionTags) DeleteByMinion(tags ...string) MinionTags {
	return nil
}

// ReplaceAllSystemTags 节点新增或上线时，系统要刷新节点的系统标签。
// 因为每次上线节点的 goos goarch inet 可能与上次不同了。
func (mts MinionTags) ReplaceAllSystemTags(newSystemTags ...string) MinionTags {
	uniq := make(map[string]struct{}, 8)
	results := make(MinionTags, 0, len(mts))
	for _, tag := range newSystemTags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		if _, ok := uniq[tag]; ok {
			continue
		}

		uniq[tag] = struct{}{}
		results = append(results, &MinionTag{Name: tag, Type: MinionTagSystem})
	}

	for _, mt := range mts {
		tag := strings.TrimSpace(mt.Name)
		if tag == "" {
			continue
		}
		if _, ok := uniq[tag]; ok {
			continue
		}
		uniq[tag] = struct{}{}
		results = append(results, &MinionTag{Name: tag, Type: mt.Type})
	}

	return results
}
