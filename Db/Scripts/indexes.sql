---------- attendance ----------

CREATE INDEX idx_attendance_student_date ON attendance (student_id, date);
CREATE INDEX idx_attendance_course_plan_date ON attendance (course_plan_id, date);