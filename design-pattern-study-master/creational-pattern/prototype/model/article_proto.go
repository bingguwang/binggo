package model

import (
    "encoding/json"
    "time"
)

type Article struct {
    Title     string
    Likes     int
    UpdatedAt time.Time
    Desc      *Desc
}
type Desc struct {
    Count int
}

// Clone 进行深拷贝，使用反序列化和序列化实现
func (a *Article) Clone() ProtoInterface {
    marshal, _ := json.Marshal(a)
    res := Article{}
    _ = json.Unmarshal(marshal, &res)
    return &res
}

type Articles map[string]*Article

func (as *Articles) Clone() Articles {
    newArticles := Articles{}

    // 替换掉需要更新的字段，这里用的是深拷贝
    for k, a := range *as {
        art := a.Clone().(*Article)
        art.Title = art.Title + "_clone"
        newArticles[k] = art
    }
    return newArticles
}
