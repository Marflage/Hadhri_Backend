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

CREATE TRIGGER update_student_enrollments_updated_at
    BEFORE UPDATE
    ON student_enrollments
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();