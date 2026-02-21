---------- courses ----------

INSERT INTO courses(name)
VALUES ('Quranic Sciences');

INSERT INTO courses(name)
VALUES ('Quran and Hadeeth Dimensions');

INSERT INTO courses(name)
VALUES ('Dars e Nizami');

---------- class_schedules ----------

INSERT INTO class_schedules(name)
VALUES ('Weekday');

INSERT INTO class_schedules(name)
VALUES ('Weekend');

---------- class_sessions ----------

INSERT INTO class_sessions(name, start_time, end_time)
VALUES ('Morning', '07:00', '11:00');

INSERT INTO class_sessions(name, start_time, end_time)
VALUES ('Afternoon', '13:30', '17:30');

INSERT INTO class_sessions(name, start_time, end_time)
VALUES ('Evening', '19:00', '21:30');

---------- course_plans ----------

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (1, 1, 3, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (1, 2, 1, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (1, 2, 2, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (1, 2, 3, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (2, 1, 1, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (2, 1, 3, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (2, 2, 1, TRUE);

INSERT INTO course_plans(course_id, class_schedule_id, class_session_id, is_active)
VALUES (2, 2, 2, TRUE);