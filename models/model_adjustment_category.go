package models

type MAdjustmentCategory struct {
	ID    int64  `json:"id" gorm:"primaryKey;autoIncrement:true" example:"1"`
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"20 days extra"`
}
type MAdjustmentCategories []*MAdjustmentCategory

var AdjustmentCategories = MAdjustmentCategories{
	{ID: AdjustmentCategoryIDUnknown, Title: "Unknown"},
	{ID: AdjustmentCategoryIDMariage, Title: "تاهل"},
	{ID: AdjustmentCategoryIDFirstChild, Title: "فرزند اول"},
	{ID: AdjustmentCategoryIDSecondChild, Title: "فرزند دوم"},
	{ID: AdjustmentCategoryIDThirdChild, Title: "فرزند سوم"},
	{ID: AdjustmentCategoryIDFourthChild, Title: "فرزند چهارم"},
	{ID: AdjustmentCategoryServiceAtOtherUnit, Title: "خدمت در یگان دیگر"},
	{ID: AdjustmentCategoryIDFatherVeteranity, Title: "ایثارگری یا جانبازی پدر"},
	{ID: AdjustmentCategoryIDMission, Title: "ماموریت"},
	{ID: AdjustmentCategoryIDCourtOrder, Title: "حکم دادگاه"},
	{ID: AdjustmentCategoryIDAbsence, Title: "غیبت"},
	{ID: AdjustmentCategoryTooMuchLeave, Title: "مرخصی بیش از حد مجاز"},
	{ID: AdjustmentCategoryIDOther, Title: "سایر"},
}

const (
	AdjustmentCategoryIDUnknown = iota
	AdjustmentCategoryIDMariage
	AdjustmentCategoryIDFirstChild
	AdjustmentCategoryIDSecondChild
	AdjustmentCategoryIDThirdChild
	AdjustmentCategoryIDFourthChild
	AdjustmentCategoryServiceAtOtherUnit
	AdjustmentCategoryIDFatherVeteranity
	AdjustmentCategoryIDMission

	AdjustmentCategoryIDCourtOrder
	AdjustmentCategoryIDAbsence
	AdjustmentCategoryTooMuchLeave
	AdjustmentCategoryIDOther
)

func (m MAdjustmentCategories) GetByID(id int64) *MAdjustmentCategory {
	for _, category := range m {
		if category.ID == id {
			return category
		}
	}
	return m[AdjustmentCategoryIDUnknown]
}
