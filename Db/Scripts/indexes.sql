---------- attendance ----------

CREATE INDEX idx_attendance_student_date ON attendance (student_id, date);
CREATE INDEX idx_attendance_course_plan_date ON attendance (course_plan_id, date);

---------- leave_requests ----------

CREATE INDEX idx_leave_requests_student_dates
    ON leave_requests (student_id, start_date DESC);

CREATE INDEX idx_leave_requests_admin_pending
    ON leave_requests (status, inserted_at ASC)
    WHERE status = 'pending';

CREATE INDEX idx_leave_requests_overlap_check
    ON leave_requests (student_id, start_date, end_date)
    WHERE status IN ('pending', 'approved');

---------- sign_up_requests ----------

CREATE INDEX idx_sign_up_requests_status
ON sign_up_requests (status)
WHERE status = 'pending';