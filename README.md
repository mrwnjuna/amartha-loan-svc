# Amartha Loan Service API
This repository contains the documentation for the Amartha Loan Service API, which allows you to create, approve, fund, and disburse loans.

## Flow
![image](https://github.com/user-attachments/assets/e4c95858-56f2-4fb1-a498-f6a490ec2f2e)

## DB Design
![loan_service_amartha_db_design](https://github.com/user-attachments/assets/21c281eb-3d6e-4338-af5e-1408e7adcd32)

## Base URL

```
https://amartha-loan-svc-8bc82f694dbf.herokuapp.com/amartha
```

---

## Endpoints

### 1. Create Loan

**Endpoint**: `/create-loan`  
**Method**: `POST`  
**Description**: Creates a new loan.

#### Request
```bash
curl --location 'https://amartha-loan-svc-8bc82f694dbf.herokuapp.com/amartha/create-loan' \
--header 'Content-Type: application/json' \
--data '{
    "borrower_id": 6,
    "principal_amount": 1000000,
    "rate": 10.5,
    "roi": 10,
    "agreement_letter": "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
}'
```

#### Response
```json
{
    "loan_id": 9
}
```

---

### 2. Approve Loan

**Endpoint**: `/approve/{loan_id}`  
**Method**: `POST`  
**Description**: Approves a loan.

#### Request
```bash
curl --location 'https://amartha-loan-svc-8bc82f694dbf.herokuapp.com/amartha/approve/9' \
--header 'Content-Type: application/json' \
--data '{
    "visit_doc_proof": "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
    "field_validator_id": "001"
}'
```

#### Response
```json
{
    "message": "Successfully approved loan",
    "status": 200
}
```

---

### 3. Fund Loan

**Endpoint**: `/fund`  
**Method**: `POST`  
**Description**: Funds a loan.

#### Request
```bash
curl --location 'https://amartha-loan-svc-8bc82f694dbf.herokuapp.com/amartha/fund' \
--header 'Content-Type: application/json' \
--data '{
    "loan_id": 9,
    "lender_id": 7,
    "funding_amount": 1000000
}'
```

#### Response
```json
{
    "message": "Successfully funded loan",
    "status": 200
}
```

#### Example email service
<img width="624" alt="Screenshot 2024-12-23 at 19 33 30" src="https://github.com/user-attachments/assets/72557a38-7ed7-409f-81bc-5be3050d6c9c" />

---

### 4. Disburse Loan

**Endpoint**: `/disburse`  
**Method**: `POST`  
**Description**: Disburses a loan.

#### Request
```bash
curl --location 'https://amartha-loan-svc-8bc82f694dbf.herokuapp.com/amartha/disburse' \
--header 'Content-Type: application/json' \
--data '{
    "loan_id": 6,
    "signed_agreement_letter": "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
    "collector_id": "002"
}'
```

#### Response
```json
{
    "message": "Successfully disbursed loan",
    "status": 200
}
```

---
