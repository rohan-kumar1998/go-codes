package main

type Tabler interface {
	TableName() string
}

type Searcher interface {
	GetID() int
}

//--------------------------------

type Syllabus struct {
	Id       int       `gorm:"primary_key" json:"id"`
	Name     string    `json:"name"`
	IsActive bool      `json:"is_active,omitempty"`
	Subjects []Subject `gorm:"foreignKey:SylId" json:"subjects,omitempty"`
}

func (Syllabus) TableName() string {
	return "syllabus"
}

//--------------------------------

type Concept struct {
	Id         int    `gorm:"primary_key" json:"id"`
	SubtopicId int    `json:"subtopic_id"`
	Name       string `json:"name"`
	IsActive   bool   `json:"is_active,omitempty"`
}

func (Concept) TableName() string {
	return "Concept"
}

//--------------------------------
type ConceptQuestion struct {
	ConceptId  int `gorm:"primary_key" json:"concept_id"`
	QuestionId int `gorm:"primary_key" json:"question_id"`
}

func (ConceptQuestion) TableName() string {
	return "concept_question"
}

//--------------------------------
type Journey struct {
	Id       int    `gorm:"primary_key" json:"id"`
	TopId    int    `json:"topic_id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active,omitempty"`
}

func (Journey) TableName() string {
	return "journey"
}

//--------------------------------
type Node struct {
	Id           int    `gorm:"primary_key" json:"id"`
	JourneyId    int    `json:"journey_id"`
	QuestionId   int    `json:"question_id"`
	VidSegmentId int    `gorm:"column:vidSegment_id" json:"videoSegment_id"`
	Name         string `json:"name"`
	IsActive     bool   `json:"is_active,omitempty"`
}

func (Node) TableName() string {
	return "Node"
}

//--------------------------------
type Question struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Answer   string `json:"answer"`
	Choice1  string `json:"choice1"`
	Choice2  string `json:"choice2"`
	Choice3  string `json:"choice3"`
	IsActive bool   `json:"is_active,omitempty"`
}

func (Question) TableName() string {
	return "Question"
}

//--------------------------------
type Subject struct {
	Id       int     `gorm:"primary_key" json:"id"`
	SylId    int     `json:"syl_id"`
	Name     string  `json:"name"`
	IsActive bool    `json:"is_active,omitempty"`
	Topics   []Topic `gorm:"foreignKey:SubId" json:"topics,omitempty"`
}

func (Subject) TableName() string {
	return "subject"
}

//--------------------------------
type Subtopic struct {
	Id       int    `gorm:"primary_key" json:"id"`
	TopId    int    `json:"topic_id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active,omitempty"`
}

func (Subtopic) TableName() string {
	return "subtopic"
}

//--------------------------------
type Topic struct {
	Id        int        `gorm:"primary_key" json:"id"`
	SubId     int        `json:"sub_id"`
	Name      string     `json:"name"`
	IsActive  bool       `json:"is_active,omitempty"`
	Subtopics []Subtopic `gorm:"foreignKey:TopId" json:"subtopics"`
}

func (Topic) TableName() string {
	return "Topic"
}

//--------------------------------
type Video struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active,omitempty"`
}

func (Video) TableName() string {
	return "Video"
}

//--------------------------------
type VideoSegment struct {
	Id        int    `gorm:"primary_key" json:"id"`
	ConceptId int    `json:"concept_id"`
	VideoId   int    `json:"video_id"`
	Name      string `json:"name"`
	VideoName string `json:"video_name"`
	IsActive  bool   `json:"is_active,omitempty"`
}

func (VideoSegment) TableName() string {
	return "Video_segment"
}

//--------------------------------
