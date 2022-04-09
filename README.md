# Author
Name: Arief Syaifuddin
Class: Golang
Link Video: 
Link Docs (Swagger UI): https://get-nearby-velezer.herokuapp.com/docs/

# Project Requirements
1. Generate Places: You can find the code in https://github.com/Velezer/get-nearby-places-jabar/blob/main/models/place.go
2. Search API: You can find the docs in https://get-nearby-velezer.herokuapp.com/docs/
3. Repository: This is the repository for the project submission

# NOTES
I use postgres to save the places data. It has 104.668 rows (places) + 9 rows (categories).
Due to limitation of free resources, query will be slow.

To handle slow query, I have implemented simple in memory server side caching with ttl (time to live).
When you run the query it will be cached for 5 minutes.
 
