package models

import "time"

type Company struct {
	// Unique Id
	Id       			int
	// The title of the company
	Title				string
	// Company logo, not necessary
	Logo	 			string
	// URL address on linkmoo.com/companies/*
	PublicURL			string
	// Company official website or service website
	Website				string
	// Which industry does the company belongs to
	Industry	 		int
	// A deeper subdividing of the market that company belongs to
	Subdivide			int
	// A approximation of the company size,
	// classified in to multi-levels by employee number
	ApproximateSize		int
	// The type of the company,
	// how it forms or how it works
	CompanyType			int
	// The pre-ipo or ipo status
	Status				int
	// A small tagline, not necessary
	Tagline				string
	// A longer description, not necessary but highly recommend
	Desc				string
	// Count the number of task that current avalibale to do
	TaskPublishing		int
	// Count how many users link to think company on linkmoo.com in any forms
	JoinedMemberNum		int
	// The company's status on linkmoo.com
	MooStatus			int
	// When does the company join linkmoo.com
	JoinTime 			time.Time	`orm:"auto_now_add;type(datetime)"`
	// The funding date of the company
	FundingTime			time.Time

	// ========== One to One ==========

	// ========== One to Many ==========

	// ========== Many to many ==========

	// ========== Reverse relationships ==========
}