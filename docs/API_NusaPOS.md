# Spec. API - NusaPOS

## 1. Authentication

### 1.1. Login

| URL                  | HTTP Method | Content Type     |
| -------------------- | ----------- | ---------------- |
| {baseUrl}/auth/login | `POST`      | application/json |

```json
// request body
{
  "username": "john.doe",
  "password": "12345678"
}
```

```json
// success response
{
  "access_token": "<access-token>",
  "token_type": "bearer",
  "expires_in": 900
}
```

## 2. Branch Management

### 2.1. Add New Branch

| URL                | HTTP Method | Content Type     |
| ------------------ | ----------- | ---------------- |
| {baseUrl}/branches | `POST`      | application/json |

Role Access: admin

```json
// request body
{
  "name": "Toko A - Cabang Mampang Prapatan",
  "address": "Mampang Prapatan, Jakarta Selatan"
}
```

```json
// success response
{
  "success": true,
  "message": "Cabang berhasil ditambahkan"
}
```

### 2.2. Get Branches

| URL                | HTTP Method | Content Type     |
| ------------------ | ----------- | ---------------- |
| {baseUrl}/branches | `GET`       | application/json |

Query Params: page, size, search

Role Access: admin

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil data cabang",
  "data": [
    {
      "id": "<uuid>",
      "name": "Toko A - Cabang Mampang Prapatan",
      "address": "Mampang Prapatan, Jakarta Selatan"
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "page_count": 1
  }
}
```

### 2.3 Edit Branch

| URL                     | HTTP Method | Content Type     |
| ----------------------- | ----------- | ---------------- |
| {baseUrl}/branches/{id} | `PUT`       | application/json |

Role Access: admin

```json
// request body
{
  "name": "Toko Mampang Prapatan",
  "address": "Mampang Prapatan, Jakarta Selatan"
}
```

```json
// success response
{
  "success": true,
  "message": "Data cabang berhasil diubah"
}
```

### 2.4. Delete Branch

| URL                     | HTTP Method | Content Type     |
| ----------------------- | ----------- | ---------------- |
| {baseUrl}/branches/{id} | `DELETE`    | application/json |

Role Access: admin

```json
// success response
{
  "success": true,
  "message": "Cabang berhasil dihapus"
}
```

## 3. Category Management

### 3.1. Add New Category

| URL                     | HTTP Method | Content Type     |
| ----------------------- | ----------- | ---------------- |
| {baseUrl}/categories    | `POST`      | application/json |

Role Access: admin

```json
// request body
{
  "name": "Coffee"
}
```

```json
// success response
{
  "success": true,
  "message": "Kategori berhasil ditambahkan"
}
```

### 3.2. Get Categories

| URL                  | HTTP Method | Content Type     |
| -------------------- | ----------- | ---------------- |
| {baseUrl}/categories | `GET`       | application/json |

Query Params: pageable, page, size, search

Role Access: manager, admin, cashier

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil data kategori",
  "data": [
    {
      "id": "<uuid>",
      "name": "Coffee"
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "page_count": 1
  }
}
```

### 3.3. Edit Category

| URL                       | HTTP Method | Content Type     |
| ------------------------- | ----------- | ---------------- |
| {baseUrl}/categories/{id} | `PUT`       | application/json |

Role Access: admin

```json
// request body
{
  "name": "Smooties"
}
```

```json
// success response
{
  "success": true,
  "message": "Data kategori berhasil diubah"
}
```

### 3.4. Delete Category

| URL                     | HTTP Method   | Content Type     |
| ----------------------- | ------------- | ---------------- |
| {baseUrl}/categories/{id} | `DELETE`    | application/json |

Role Access: admin

```json
// success response
{
  "success": true,
  "message": "Kategori berhasil dihapus"
}
```

## 4. Product Management

### 4.1. Add New Product

| URL                   | HTTP Method | Content Type     |
| --------------------- | ----------- | ---------------- |
| {baseUrl}/products    | `POST`      | application/json |

Role Access: manager, admin

```json
// request body
{
  "category_id": "<uuid>",
  "branch_id": "<uuid>",
  "name": "Kopi Susu Keluarga",
  "price": 16000,
  "is_available": true
}
```

```json
// success response
{
  "success": true,
  "message": "Produk berhasil ditambahkan"
}
```

### 4.2. Get Products

| URL                   | HTTP Method | Content Type     |
| --------------------- | ----------- | ---------------- |
| {baseUrl}/products    | `GET`       | application/json |

Query Params: pageable, is_available, category_id, page, size, search

Role Access: manager, admin, cashier

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil data produk",
  "data": [
    {
      "id": "<uuid>",
      "category_id": "<uuid>",
      "category_name": "Coffee",
      "branch_id": "<uuid>",
      "branch_name": "Toko Mampang Prapatan",
      "name": "Kopi Susu Keluarga",
      "price": 16000,
      "is_available": true
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "page_count": 1
  }
}
```

### 4.3. Edit Product

| URL                        | HTTP Method | Content Type     |
| -------------------------- | ----------- | ---------------- |
| {baseUrl}/products/{id}    | `PUT`       | application/json |

Role Access: manager, admin

```json
// request body
{
  "category_id": "<uuid>",
  "branch_id": "<uuid>",
  "name": "Butterscotch",
  "price": 17000,
  "is_available": true
}
```

```json
// success response
{
  "success": true,
  "message": "Data produk berhasil diubah"
}
```

### 4.4. Delete Product

| URL                        | HTTP Method | Content Type     |
| -------------------------- | ----------- | ---------------- |
| {baseUrl}/products/{id}    | `DELETE`    | application/json |

Role Access: manager, admin

```json
// success response
{
  "success": true,
  "message": "produk berhasil dihapus"
}
```

## 5. Order Management

### 5.1. Place Order

| URL                 | HTTP Method | Content Type     |
| ------------------- | ----------- | ---------------- |
| {baseUrl}/orders    | `POST`      | application/json |

Role Access: admin, cashier

```json
// request body
{
  "branch_id": "<uuid>",
  "payment_method": "cash",
  "paid_amount": 20000,
  "details": [
    {
      "product_id": "<uuid>",
      "quantity": 1,
      "price": 16000
    }
  ]
}
```

```json
// success response
{
  "success": true,
  "message": "Pesanan berhasil dibuat",
  "data": {
    "branch_id": "<uuid>",
    "branch_name": "Toko Mampang Prapatan",
    "cashier_id": "<uuid>",
    "cashier_name": "John Doe",
    "total_amount": 16000,
    "payment_method": "cash",
    "paid_amount": 20000,
    "change_amount": 4000,
    "order_date": "2025-08-13 15:53",
    "details": [
      {
        "product_id": "<uuid>",
        "product_name": "Kopi Susu Keluarga",
        "quantity": 1,
        "price": 16000,
        "subtotal": 16000
      }
    ]
  }
}
```

## 6. Sales Reporting

### 6.1. Transaction Detail

| URL                          | HTTP Method | Content Type     |
| ---------------------------- | ----------- | ---------------- |
| {baseUrl}/transaction/detail | `GET`       | application/json |

Query Params: page, size, start_date, end_date, cashier_id, payment_method

Role Access: admin, manager

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil data transaksi detil",
  "data": [
    {
      "id": "<uuid>",
      "branch_id": "<uuid>",
      "branch_name": "Toko Mampang Prapatan",
      "cashier_id": "<uuid>",
      "cashier_name": "John Doe",
      "total_amount": 16000,
      "payment_method": "cash",
      "order_date": "2025-08-13 15:53"
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "page_count": 1
  }
}
```

### 6.2. Summary Transaction Daily

| URL                                 | HTTP Method | Content Type     |
| ----------------------------------- | ----------- | ---------------- |
| {baseUrl}/transaction/summary/daily | `GET`       | application/json |

Query Params: month, year

Role Access: admin, manager

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil rekap transaksi harian",
  "data": [
    {
      "order_date": "2025-08-01",
      "total": 100000
    }
  ]
}
```

### 6.3. Summary Transaction Weekly

| URL                                  | HTTP Method | Content Type     |
| ------------------------------------ | ----------- | ---------------- |
| {baseUrl}/transaction/summary/weekly | `GET`       | application/json |

Query Params: month, year

Role Access: admin, manager

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil rekap transaksi mingguan",
  "data": [
    {
      "week": "Minggu ke-1",
      "total": 400000
    }
  ]
}
```

### 6.4. Summary Transaction Monthly

| URL                                   | HTTP Method | Content Type     |
| ------------------------------------- | ----------- | ---------------- |
| {baseUrl}/transaction/summary/monthly | `GET`       | application/json |

Query Params: year

Role Access: admin, manager

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil rekap transaksi bulanan",
  "data": [
    {
      "month": "Januari",
      "total": 1400000
    }
  ]
}
```

### 6.5. Top Selling Products

| URL                                | HTTP Method | Content Type     |
| ---------------------------------- | ----------- | ---------------- |
| {baseUrl}/transaction/top/products | `GET`       | application/json |

Role Access: manager, admin, cashier

```json
// success response
{
  "success": true,
  "message": "Berhasil mengambil data produk terlaris",
  "data": [
    {
      "product_name": "Kopi Susu Keluarga",
      "total": 157
    }
  ]
}
```

## 7. User Management

### 7.1. Create New User

| URL                | HTTP Method | Content Type     |
| ------------------ | ----------- | ---------------- |
| {baseUrl}/users    | `POST`      | application/json |

Role Access: admin

```json
// request body
{
  "name": "John Doe",
  "username": "john.doe",
  "password": "12345678",
  "role": "cashier"
}
```

```json
// success response
{
  "success": true,
  "message": "User berhasil dibuat"
}
```

### 7.2. Get Users

| URL                | HTTP Method | Content Type     |
| ------------------ | ----------- | ---------------- |
| {baseUrl}/users    | `GET`       | application/json |

Query Params: page, size, search

Role Access: admin

```json
{
  "success": true,
  "message": "Berhasil mengambil data user",
  "data": [
    {
      "id": "<uuid>",
      "name": "John Doe",
      "username": "john.doe",
      "role": "cashier"
    }
  ],
  "meta": {
    "total": 1,
    "page": 1,
    "page_count": 1
  }
}
```

### 7.3. Edit User

| URL                  | HTTP Method | Content Type     |
| -------------------- | ----------- | ---------------- |
| {baseUrl}/users/{id} | `PUT`       | application/json |

Role Access: admin

```json
// request body
{
  "name": "Jane Doe",
  "username": "jane.doe",
  "role": "manager"
}
```

```json
// success response
{
  "success": true,
  "message": "Data user berhasil diubah"
}
```

### 7.4. Delete User

| URL                  | HTTP Method    | Content Type     |
| -------------------- | -------------- | ---------------- |
| {baseUrl}/users/{id} | `DELETE`       | application/json |

Role Access: admin

```json
// success response
{
  "success": true,
  "message": "User berhasil dihapus"
}
```