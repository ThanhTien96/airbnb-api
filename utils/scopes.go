package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ThanhTien96/airbnb-api/internal/common"

	"gorm.io/gorm"
)

func GetOrdernPattern(sorts []string) string {
	var orderPattern string
	for _, sort := range sorts {
		sortType := "asc"
		if strings.HasPrefix(sort, "-") {
			sort = sort[1:]
			sortType = "desc"
		}
		orderPattern += fmt.Sprintf("%s %s, ", sort, sortType)
	}
	orderPattern = strings.TrimRight(orderPattern, ", ")
	return orderPattern
}

func GetRangePattern(field string) string {
	return field + " >= ? and " + field + " <= ?"
}

func SearchPattern(field string) string {
	return field + " ILIKE ?"
}

func ApplyFilterQuery(filterParams *common.FilterParams, db *gorm.DB) (*gorm.DB, *common.Paging, error) {
	var paging common.Paging
	paging.Process()
	// Add sorting
	if len(filterParams.Sorts) != 0 {
		orderPattern := GetOrdernPattern(filterParams.Sorts)
		db = db.Order(orderPattern)
	}
	// Add limit
	if filterParams.Limit == "" {
		filterParams.Limit = fmt.Sprint(paging.Limit)
	}
	limit, err := strconv.Atoi(filterParams.Limit)
	paging.Limit = limit
	if err != nil {
		return nil, nil, err
	}
	db = db.Limit(limit)
	// Add offset
	if filterParams.Page != "" {
		page, err := strconv.Atoi(filterParams.Page)
		paging.Page = page
		offset := (page - 1) * paging.Limit
		if err != nil {
			return nil, nil, err
		}
		db = db.Offset(offset)
	}
	// Add range filter
	if filterParams.StartDate != "" && filterParams.EndDate != "" {
		startDateInt, err := strconv.ParseInt(filterParams.StartDate, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		endDateInt, err := strconv.ParseInt(filterParams.EndDate, 10, 64)
		if err != nil {
			return nil, nil, err
		}
		queryString := GetRangePattern("created_at")

		db = db.Where(queryString, startDateInt, endDateInt)

	}

	// search
	if filterParams.Search != "" {
		queryStr := SearchPattern("mov_title")
		fmt.Println("qe ", queryStr)

		query := filterParams.Search + "%"

		db = db.Debug().Where(queryStr, query)

	}
	return db, &paging, nil
}
