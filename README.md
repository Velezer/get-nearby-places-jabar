# Author
Name: Arief Syaifuddin
Class: Golang
Link Video: 
Link Docs (Swagger UI): https://get-nearby-velezer.herokuapp.com/docs/

# NOTES
I use postgres to save the places data. It has 104.668 rows (places) + 9 rows (categories).
Due to limitation of free resources, query will be slow.

To handle slow query, I have implemented simple in memory server side caching with ttl (time to live).
When you run the query it will be cached for 5 minutes.
 
