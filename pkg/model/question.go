package model

type Question struct {
	ID                           int                `db:"id,pk" json:"id"`
	AuthorID                     int                `db:"author_id" json:"author_id"`
	AuthorRole                   AuthorRole         `db:"author_role" json:"author_role"`
	Status                       QuestionStatus     `db:"status" json:"status"`
	QuestionType                 QuestionType       `db:"question_type" json:"question_type"`
	DataType                     DataType           `db:"data_type" json:"data_type"`
	QuestionValue                string             `db:"question_value" json:"question_value"`
	QuestionLabel                string             `db:"question_label" json:"question_label"`
	PromptText                   string             `json:"prompt_text"`
	RequiredAnswer               bool               `json:"required_answer"`
	AllowMultipleAnswer          bool               `json:"allow_multiple_answer"`
	HistoricallyTrackedParameter bool               `json:"historically_tracked_parameter"`
	ParentQuestionID             int                `json:"parent_question_id"`
	AnswerValues                 []string           `json:"answer_values"`
	MultipleSelection            bool               `json:"multiple_selection"`
	MinAdditions                 float64            `json:"min_additions"`
	MaxAdditions                 float64            `json:"max_additions"`
	MinSymbols                   int                `json:"min_symbols"`
	MaxSymbols                   int                `json:"max_symbols"`
	RowsNumber                   int                `json:"rows_number"`
	KeywordsRecognize            bool               `json:"keywords_recognize"`
	DateType                     DateType           `json:"date_type"`
	AllowCurrentAndFutureDate    bool               `json:"allow_current_and_future_date"`
	AllowApproximateDate         bool               `json:"allow_approximate_date"`
	MeasurementUnits             []MeasurementUnit  `json:"measurement_units"`
	DefaultMeasurementUnit       MeasurementUnit    `json:"default_measurement_unit"`
	MaxValue                     float64            `json:"max_value"`
	MinValue                     float64            `json:"min_value"`
	AttachmentFormats            []AttachmentFormat `json:"attachment_formats"`
	CreatedTable                 Table              `json:"created_table"`
	RefBook                      RefBook            `json:"ref_book"`
	AllowOther                   bool               `json:"allow_other"`
	MedDataType                  MedDataType        `json:"med_data_type"`
	Biomarker                    Biomarker          `json:"biomarker"`
	//HealthStatusImpactAnswer     interface{}        `json:"health_status_impact_answer"` //todo
	AdditionalQuestionInfluence bool      `json:"additional_question_influence"`
	DefaultChapter              Chapter   `json:"default_chapter"`
	DefaultBlock                Block     `json:"default_block"`
	DestinationChapters         []Chapter `json:"destination_chapters"`
	DestinationBlocks           []Block   `json:"destination_blocks"`
	AgeGroups                   AgeGroup  `json:"age_groups"`
	Genders                     []Gender  `json:"genders"`
	Comment                     string    `json:"comment"`
	Notification                string    `json:"notification"`
	//AdditionalQuestions          interface{}        `json:"additional_questions"` //todo
	//And                          interface{}        `json:"and"`                  //todo
	Divider bool `json:"divider"`
}
