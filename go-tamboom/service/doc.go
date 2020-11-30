// Package service contains functions that handle the main business logic
//
// service handles file processing, as well as holding structs containing the full service code.
// The business logic handles the process like a batch job; this is a process meant to be handled during low
// usage hours. The reasoning for this is the nature of handling a large file and processing it makes it seem to be a batch job
//
// Optimization can still be done, but things like usage of asynchronous functions or goroutines were not prioritized over code correctness, and code organization.
// The reason for this is that batch jobs are usually approached with a sense of utilizing resources during off-hours for a business, allowing things
// to take the time to get things processed correctly
//
// summary handles process summarization, handling the total amount of donations, and failed donations
// Data for the summary, such as average donations, and top donors, are based only on successful donations.
// The reasoning for this is that it does not make sense to handle data for failed or invalid donations and to use them
// as a measure for the total overall (i.e. A top donor should never be from a failed donation)
package service
