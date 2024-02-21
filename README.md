***Instructions***
1) Execute the R_ABC.R script with the command "Rscript R_ABC.R".  In the terminal it will ouput the estimated Mu value and the execution time it took to run the script and provide the mu.

2) Execute the main.go script with the command go run "main.go".  In the termimal it will output the estimated Mu value and the execution time it took to run the script. 

3) Execute the main_test.go script with the command go run "go test".  The script will check to see if the mu value from the main.go script is within .25 of the Mu value produced by the R script. 

4) Execute the benchmarking using the command "go test -bench=BenchmarkSimulateData."  

***Notes:***  

I attempted to perform the comparison between R and Go using several different methods outlined in the Gelman and Vehtari article.  For example, I originally attempted to build an overparameterized model with regulation.  Specifically, I couldn't find a Go library that used Ridge or Lasso.  When I removed regulation I continued to get different results.  Accordingly, I simplified a bit and moved to using a generic algorthm.  Specifically, I went with an Approximate Baysesian Computation (ABC) as there was a way to do this in both R and Go.  

Overall, I found the performance of the Rscript to be faster than Go.  I didn't even need to use any specific R libraries in order to generate the ABC.  Rather, I was able to simply leverage commands defined within base R.  R was just way faster than Go.  I initially attempted to optimize Go using concurrency and channels.  The issue there was that I was unable to use seeding inside of concurrency.  Since I kept getting a different result, I killed the concurrency approach.  

That said, I did implement some optimizations in the main.go script.  Specifically, I preallocated slices since I knew the upper limit. Using preallocated slices translated to reducing the time spent on memory allocation. Overall, the version of my main.go file with preallocated faster than the version without.  

I found the benchmarking test to be useful.  The metadata it provided about my laptop was interesting and useful to know that it allowed Go to use up to 16 operating threads.  The abilty to isolate a specific function was also helpful and understanding the performance time of that segment was useful.

***Recommendation to Research Consultancy***
Overall, I would not recommend shifting from R to Go globally.  Rather, I would entertain using Go in unique circumstances where performance may be a concern in R.  With this test I found R to be superior to Go in two ways:  1) It had a better set of mature, statistically oriented libraries.  In Go, I couldn't even find a library that I could use for regularization.  2) R performed, on average, 10X faster than I was able to get Go to perform. 

Because R was running much faster than Go I don't believe there were any cloud computing cost benefits of leveraging Go over R.  (At least based on the scripts I developed.)



