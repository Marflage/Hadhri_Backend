CREATE TRIGGER update_courses_updated_at
    BEFORE UPDATE
    ON courses
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_class_schedules_updated_at
    BEFORE UPDATE
    ON class_schedules
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_class_sessions_updated_at
    BEFORE UPDATE
    ON class_sessions
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_course_plans_updated_at
    BEFORE UPDATE
    ON course_plans
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_available_semesters_updated_at
    BEFORE UPDATE
    ON available_semesters
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_students_updated_at
    BEFORE UPDATE
    ON students
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_enrollments_updated_at
    BEFORE UPDATE
    ON enrollments
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_admins_updated_at
    BEFORE UPDATE
    ON admins
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER admin_insert_protection
    BEFORE INSERT
    ON admins
    FOR EACH ROW
EXECUTE FUNCTION prevent_admin_insert();    

CREATE TRIGGER update_attendance_statuses_updated_at
    BEFORE UPDATE
    ON attendance_statuses
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_attendance_updated_at
    BEFORE UPDATE
    ON attendance
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_account_activation_requests_updated_at
    BEFORE UPDATE
    ON account_activation_requests
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_account_activation_requests_status_changed_at
    BEFORE UPDATE
    ON account_activation_requests
    FOR EACH ROW
EXECUTE FUNCTION update_status_changed_at_column();