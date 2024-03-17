CREATE TABLE IF NOT EXISTS "patients" (
    "id" VARCHAR(200) NOT NULL,
    "name" VARCHAR(200),
    "email" VARCHAR(200),
    "phone" VARCHAR(20),
    "dob" VARCHAR(200),
    "age" VARCHAR(200),
    "city" VARCHAR(200),
    "state" VARCHAR(200),
    "district" VARCHAR(200),
    "block" VARCHAR(200),
    "pincode" VARCHAR(200),
    "gender" VARCHAR(200),
    "aadhar_number" VARCHAR(200),
    "relation_name" VARCHAR(200),
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "doctors" (
    "id" VARCHAR(200) NOT NULL,
    "name" VARCHAR(200),
    "dob" VARCHAR(200),
    "age" VARCHAR(200),
    "city" VARCHAR(200),
    "state" VARCHAR(200),
    "district" VARCHAR(200),
    "pincode" VARCHAR(200),
    "gender" VARCHAR(200),
    "email" VARCHAR(200),
    "department" VARCHAR(200),
    "designation" VARCHAR(200),
    "license_number" VARCHAR(200),
    "phone" VARCHAR(20),
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP,
    PRIMARY KEY ("id")
    );

CREATE TABLE IF NOT EXISTS "doctor_followup" (
    "id" VARCHAR(200),
    "doctor_id" VARCHAR(200),
    "date" VARCHAR(200),
    "time" VARCHAR(200),
    "followup_count" SERIAL NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "patient_followup" (
    "id" VARCHAR(200),
    "patient_id" VARCHAR(200),
    "date" VARCHAR(200),
    "time" VARCHAR(200),
    "followup_count" SERIAL NOT NULL,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY ("id")
);

ALTER TABLE doctor_followup
ADD CONSTRAINT fk_doctor_followup
FOREIGN KEY (doctor_id) REFERENCES doctors(id);

ALTER TABLE patient_followup
ADD CONSTRAINT fk_patient_followup
FOREIGN KEY (patient_id) REFERENCES patients(id);