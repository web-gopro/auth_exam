-- users jadvali: oddiy foydalanuvchilar
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(50) NOT NULL, -- [active, deleted]
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID
);

-- sysusers: tizim yurituvchi foydalanuvchilar (admin, buxgalter va h.k.)
CREATE TABLE sysusers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(50) NOT NULL, -- [active, deleted]
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID
);

-- roles: tizim rollari (superadmin, marketolog, buxgalter va h.k.)
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(50) NOT NULL, -- [active, deleted]
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by UUID
);

-- sysuser_roles: sysuser va role orasidagi bog‘lovchi jadval (many-to-many)
CREATE TABLE sysuser_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sysuser_id UUID NOT NULL REFERENCES sysusers(id) ON DELETE CASCADE,
    role_id UUID NOT NULL REFERENCES roles(id) ON DELETE CASCADE
);


CREATE EXTENSION IF NOT EXISTS "pgcrypto";
