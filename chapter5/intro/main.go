package main

// Movie holds a movie data
type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	Boxffice  `bson:"boxOffice"`
}

// BoxOffice is nested in movie
type Boxffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}
