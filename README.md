# Mastercard
Integrating the Mastercard API with AWS &amp; Go

## Functions

Two lambda functions are implemented:
- LoadRates() gets rates from mastercard and stores in DynamoDB
- GetRates() obtains rate from DynamoDB when API Gateway is called
