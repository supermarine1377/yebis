# Yebis

## Introduction

Yebis allows you to predict a metric called "Investment Score".

The Investment Score ranges from +8 to -8.

If the Investment Score is 8, it *might* tell us that economic conditions around the world are favorable for investing in stocks, bonds, and commodities.

If the score is -8, it *might* tell us that economic conditions are not suitable for investing.

The Investment Score is calculated from several types of economic metrics, such as FEDFUNDS RATE, US10Y, and so on.

## How to use

Create .env file at the root directory in your local environment.

The file content should be like:

```
FED_API_KEY=fred_api_key
```

You can register your FED API Key here:

https://fred.stlouisfed.org/docs/api/api_key.html

