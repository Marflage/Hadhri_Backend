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
    id           SERIAL PRIMARY KEY,
    inserted_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    first_name   VARCHAR(30)         NOT NULL CHECK ( LENGTH(first_name) >= 2 ),
    last_name    VARCHAR(30)         NOT NULL CHECK ( LENGTH(last_name) >= 2 ),
    email        VARCHAR(100) UNIQUE NOT NULL CHECK ( email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    phone_number VARCHAR(11) UNIQUE  NOT NULL CHECK ( phone_number ~ '^[0-9]{11}$' ),
    password     VARCHAR(100)         NOT NULL CHECK (LENGTH(password) >= 8)
);

CREATE TABLE student_enrollments
(
    id             SERIAL PRIMARY KEY,
    inserted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    student_id     INT NOT NULL UNIQUE REFERENCES students (id),
    course_plan_id INT NOT NULL REFERENCES course_plans (id),
    enrolled_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    semester       INT NOT NULL
);

CREATE TABLE attendance
(
    id             SERIAL PRIMARY KEY,
    inserted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    update_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    student_id     INTEGER  NOT NULL REFERENCES students (id),
    course_plan_id INTEGER  NOT NULL REFERENCES course_plans (id),
    date           DATE     NOT NULL,
    status_id      SMALLINT NOT NULL REFERENCES attendance_statuses (id),
    UNIQUE (student_id, course_plan_id, date)
);