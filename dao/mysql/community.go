package mysql

import (
	"bluebell/global"
	"bluebell/models"
	"database/sql"
	"errors"
)

// GetCommunityList 得到社区列表
func GetCommunityList() (*[]models.Community, error) {
	var communities []models.Community // 直接使用切片，而不是指针
	sqlStr := "SELECT community_id, community_name FROM community"
	rows, err := global.DB.Query(sqlStr)
	if err != nil {
		return nil, err // 直接返回错误，不检查 sql.ErrNoRows
	}
	defer rows.Close()

	for rows.Next() {
		var community models.Community
		if err := rows.Scan(&community.ID, &community.Name); err != nil {
			return nil, err // 返回从 rows.Scan 接收到的错误
		}
		communities = append(communities, community) // 直接添加到切片中
	}

	if err := rows.Err(); err != nil { // 检查 rows 是否在迭代过程中产生了错误
		return nil, err
	}

	// 返回指向切片的指针（如果需要的话）
	return &communities, nil
}

// GetCommunityDetailByID 得到某个社区详情
func GetCommunityDetailByID(id int) (*models.CommunityDetail, error) {
	var cd models.CommunityDetail
	sqlStr := `SELECT community_id, community_name, introduction, create_time FROM community WHERE community_id = ?`

	err := global.DB.QueryRow(sqlStr, id).Scan(&cd.ID, &cd.Name, &cd.Introduction, &cd.CreateTime)
	if err != nil {
		if err == sql.ErrNoRows { // 区分没有数据的情况
			err = errors.New("无效ID")
			return nil, err
		}
		return nil, err
	}
	return &cd, nil
}
