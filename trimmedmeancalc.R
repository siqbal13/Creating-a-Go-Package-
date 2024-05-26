#Example verification of the results for trimmedmean calculated from go language program using R language calculations
# Read the integer data from CSV
int_data <- read.csv("int_data.csv", header = FALSE)
int_data <- as.numeric(int_data$V1)

# Read the float data from CSV
float_data <- read.csv("float_data.csv", header = FALSE)
float_data <- as.numeric(float_data$V1)

# Compute the symmetric trimmed mean with 0.05 trimming
# Again here you can use the other degrees according to your requirements as #here we are using hard code value 0.05
int_symmetric_trimmed_mean <- mean(int_data, trim = 0.05)
float_symmetric_trimmed_mean <- mean(float_data, trim = 0.05)

# Function to compute asymmetric trimmed mean
asymmetric_trimmed_mean <- function(data, lower_trim, upper_trim) {
  sorted_data <- sort(data)
  n <- length(data)
  lower_index <- floor(n * lower_trim) + 1
  upper_index <- ceiling(n * (1 - upper_trim))
  trimmed_data <- sorted_data[lower_index:upper_index]
  mean(trimmed_data)
}

# Compute the asymmetric trimmed mean with 0.05 lower trim and 0.1 upper trim

int_asymmetric_trimmed_mean <- asymmetric_trimmed_mean(int_data, 0.05, 0.1)
float_asymmetric_trimmed_mean <- asymmetric_trimmed_mean(float_data, 0.05, 0.1)

# Print the results
cat("Symmetric Trimmed Mean (0.05 trim) for integers:", int_symmetric_trimmed_mean, "\n")
cat("Symmetric Trimmed Mean (0.05 trim) for floats:", float_symmetric_trimmed_mean, "\n")
cat("Asymmetric Trimmed Mean (0.05 lower, 0.1 upper trim) for integers:", int_asymmetric_trimmed_mean, "\n")
cat("Asymmetric Trimmed Mean (0.05 lower, 0.1 upper trim) for floats:", float_asymmetric_trimmed_mean, "\n")
