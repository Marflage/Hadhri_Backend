---------- Reference Tables ----------

CREATE TABLE courses
(
    id          SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name        VARCHAR(50) UNIQUE NOT NULL CHECK ( LENGTH(name) BETWEEN 3 AND 50)
);

CREATE TABLE class_schedules
(
    id          SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name        VARCHAR(30) UNIQUE NOT NULL CHECK ( LENGTH(name) BETWEEN 3 AND 30)
);

CREATE TABLE class_sessions
(
    id          SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name        VARCHAR(10) UNIQUE NOT NULL CHECK ( LENGTH(name) BETWEEN 3 AND 10),
    start_time  TIME               NOT NULL,
    end_time    TIME               NOT NULL CHECK ( end_time > start_time )
);

CREATE TABLE course_plans
(
    id                SERIAL PRIMARY KEY,
    inserted_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    course_id         INT NOT NULL REFERENCES courses (id),
    class_schedule_id INT NOT NULL REFERENCES class_schedules (id),
    class_session_id  INT NOT NULL REFERENCES class_sessions (id),
    is_active         BOOLEAN   DEFAULT TRUE,
    UNIQUE (course_id, class_schedule_id, class_session_id)
);

CREATE TABLE available_semesters
(
    id             SERIAL PRIMARY KEY,
    inserted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    course_plan_id INT NOT NULL REFERENCES course_plans (id),
    semester       INT NOT NULL
    UNIQUE (course_plan_id, semester)
);

CREATE TABLE admins
(
    id          SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    email       VARCHAR(100) UNIQUE NOT NULL CHECK ( email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    password    VARCHAR(100)        NOT NULL CHECK ( LENGTH(password) >= 8 )
);

CREATE TABLE attendance_statuses
(
    id          SMALLSERIAL PRIMARY KEY,
    inserted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name        VARCHAR(10) NOT NULL UNIQUE
);

---------- Transactional Tables ----------

CREATE TABLE students
(
    id            INT          NOT NULL,
    inserted_at   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    full_name     VARCHAR(100) NOT NULL,
    email         citext       NOT NULL,
    phone_number  VARCHAR(11)  NOT NULL,
    password_hash VARCHAR(200) NOT NULL,

    CONSTRAINT students_pk PRIMARY KEY (id),

    CONSTRAINT students_uq_email UNIQUE (email),
    CONSTRAINT students_uq_phone_number UNIQUE (phone_number),

    CONSTRAINT students_chk_valid_id CHECK ( id > 0 ),
    CONSTRAINT students_chk_full_name CHECK ( LENGTH(full_name) >= 3 ),
    CONSTRAINT students_chk_email CHECK ( email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    CONSTRAINT students_chk_phone_number CHECK ( phone_number ~ '^[0-9]{11}$' )
);

CREATE TABLE enrollments
(
    id             INT GENERATED ALWAYS AS IDENTITY,
    inserted_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    student_id     INT         NOT NULL,
    course_plan_id INT         NOT NULL,
    enrolled_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    semester       INT         NOT NULL,

    CONSTRAINT enrollments_pk PRIMARY KEY (id),

    CONSTRAINT enrollments_student_id_unique UNIQUE (student_id),

    CONSTRAINT enrollments_student_id_fk_students_id FOREIGN KEY (student_id) REFERENCES students (id) ON DELETE RESTRICT,
    CONSTRAINT enrollments_course_plan_id_fk_course_plans_id FOREIGN KEY (course_plan_id) REFERENCES course_plans (id) ON DELETE RESTRICT
);

CREATE TABLE attendance
(
    id             SERIAL PRIMARY KEY,
    inserted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    student_id     INTEGER  NOT NULL REFERENCES students (id),
    course_plan_id INTEGER  NOT NULL REFERENCES course_plans (id),
    date           DATE     NOT NULL DEFAULT CURRENT_DATE,
    status_id      SMALLINT NOT NULL REFERENCES attendance_statuses (id),
    UNIQUE (student_id, course_plan_id, date)
);

CREATE TABLE leave_requests
(
    id                SERIAL PRIMARY KEY,
    inserted_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    student_id        INT       NOT NULL REFERENCES students (id),
    start_date        DATE      NOT NULL,
    end_date          DATE      NOT NULL,
    reason            TEXT      NOT NULL,
    status            TEXT      NOT NULL DEFAULT 'pending',
    status_changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_status CHECK ( status IN ('pending', 'approved', 'rejected', 'canceled') ),
    CONSTRAINT valid_date_range CHECK ( end_date >= start_date ),
    CONSTRAINT unique_student_leave_dates UNIQUE (student_id, start_date, end_date)
);

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE account_activation_requests
(
    id                INT GENERATED ALWAYS AS IDENTITY,
    inserted_at       TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status            VARCHAR(20)  NOT NULL DEFAULT 'pending',
    status_changed_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    full_name         VARCHAR(100) NOT NULL,
    email             citext       NOT NULL,
    phone_number      VARCHAR(11)  NOT NULL,
    password_hash     VARCHAR(200) NOT NULL,
    course_plan_id    INT          NOT NULL,
    semester          INT          NOT NULL,

    CONSTRAINT pk_id PRIMARY KEY (id),

    CONSTRAINT uq_email UNIQUE (email),
    CONSTRAINT uq_phone_number UNIQUE (phone_number),

    CONSTRAINT chk_valid_status CHECK ( status IN ('pending', 'approved', 'declined', 'canceled')),
    CONSTRAINT chk_full_name CHECK ( LENGTH(full_name) >= 3 ),
    CONSTRAINT chk_email CHECK ( email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
    CONSTRAINT chk_phone_number CHECK ( phone_number ~ '^[0-9]{11}$' ),
    CONSTRAINT chk_valid_semester CHECK ( semester > 0 ),

    CONSTRAINT fk_course_plan_id FOREIGN KEY (course_plan_id) REFERENCES course_plans (id) ON DELETE RESTRICT
);