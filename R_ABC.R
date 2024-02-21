start_time <- Sys.time() # Fire up the timer

set.seed(123) # Ensure that you get the same result.

# Generate observed data
mu_true <- 10
sigma <- 2
n <- 10000
observed_data <- rnorm(n, mean = mu_true, sd = sigma)

# ABC parameters
mu_proposals <- seq(from = 5, to = 15, by = 0.1)
tolerance <- 0.5
n_sim <- 100

# ABC algorithm
abc_results <- sapply(mu_proposals, function(mu) {
  abs(mean(rnorm(n_sim, mean = mu, sd = sigma)) - mean(observed_data)) <= tolerance
})

accepted_mu <- mu_proposals[abc_results]
estimated_mu <- mean(accepted_mu)

print(paste("Estimated mu:", estimated_mu))

end_time <- Sys.time() # Stop the timer.

# Calculate and print execution time
execution_time <- end_time - start_time
print(paste("Execution time:", execution_time))
