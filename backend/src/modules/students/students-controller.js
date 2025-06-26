const { z } = require("zod");
const { ApiError } = require("../../utils");

const {
  getAllStudents,
  addNewStudent,
  getStudentDetail,
  setStudentStatus,
  updateStudent,
} = require("./students-service");

/**
 * Type Definition for students add and create payload.
 * @typedef {Object} StudentPayload
 * @property {string} name
 * @property {string} email
 * @property {string} gender
 * @property {string} phone
 * @property {string} dob
 * @property {string} admissionDate
 * @property {string} class
 * @property {string} section
 * @property {string} roll
 * @property {string} currentAddress
 * @property {string} permanentAddress
 * @property {string} fatherName
 * @property {string} fatherPhone
 * @property {string} motherName
 * @property {string} motherPhone
 * @property {string} guardianName
 * @property {string} guardianPhone
 * @property {string} relationOfGuardian
 * @property {boolean} systemAccess
 */

/**
 * @uri /api/v1/students
 * @method GET
 * @description Get all students
 * @param {{ query: { name: string, className: string, section: string, roll: string } }} req
 */
const handleGetAllStudents = async (req, res) => {
  try {
    //  Validate request
    const payload = z.object({
      name: z.string().nullable().optional().default(null),
      className: z.string().nullable().optional().default(null),
      section: z.string().nullable().optional().default(null),
      roll: z.string().nullable().optional().default(null),
    }).parse(req.query);

    const students = await getAllStudents(payload);
    res.json({ students });
  } catch (error) {
    if (error instanceof z.ZodError) {
      throw new ApiError(400, `Validation error: ${error.errors.map(e => e.message).join(', ')}`);
    }
    throw error;
  }
};

/**
 * @uri /api/v1/students
 * @method POST
 * @description Add a new student
 * @param {{ body: StudentPayload }} req 
 */
const handleAddStudent = async (req, res) => {
  try {
    // Validate request
    const payload = z.object({
      name: z.string().min(3, "Name must be at least 3 characters").max(100, "Name must be less than 100 characters"),
      email: z.string().email("Invalid email format"),
      gender: z.enum(["Male", "Female", "Other"], { errorMap: () => ({ message: "Gender must be Male, Female, or Other" }) }),
      phone: z.string().length(10, "Phone number must be exactly 10 digits").regex(/^\d+$/, "Phone number must contain only digits"), // I have ignored the country code for now
      dob: z.string().regex(/^\d{4}-\d{2}-\d{2}$/, "Date of birth must be in YYYY-MM-DD format"),
      admissionDate: z.string().regex(/^\d{4}-\d{2}-\d{2}$/, "Admission date must be in YYYY-MM-DD format"),
      class: z.string().min(1, "Class is required").max(100, "Class name must be less than 100 characters"), // No strict validation as class name could be any string value, Ideally I would use primary key from class table
      section: z.string().min(1, "Section is required").max(100, "Section name must be less than 100 characters"),
      roll: z.coerce.number().int("Roll must be a whole number").min(1, "Roll must be greater than 0"),
      currentAddress: z.string().min(1, "Current address is required").max(100, "Current address must be less than 100 characters"),
      permanentAddress: z.string().min(1, "Permanent address is required").max(100, "Permanent address must be less than 100 characters"),
      fatherName: z.string().min(1, "Father's name is required").max(100, "Father's name must be less than 100 characters"),
      fatherPhone: z.string().length(10, "Father's phone must be exactly 10 digits").regex(/^\d+$/, "Father's phone must contain only digits"),
      motherName: z.string().min(1, "Mother's name is required").max(100, "Mother's name must be less than 100 characters"),
      motherPhone: z.string().length(10, "Mother's phone must be exactly 10 digits").regex(/^\d+$/, "Mother's phone must contain only digits"),
      guardianName: z.string().min(1, "Guardian's name is required").max(100, "Guardian's name must be less than 100 characters"),
      guardianPhone: z.string().length(10, "Guardian's phone must be exactly 10 digits").regex(/^\d+$/, "Guardian's phone must contain only digits"),
      relationOfGuardian: z.string().min(1, "Relation of guardian is required").max(100, "Relation of guardian must be less than 100 characters"),
      systemAccess: z.boolean().optional().default(false)
    }).parse(req.body);

    // Add student
    const message = await addNewStudent(payload);
    res.json(message);
  } catch (error) {
    if (error instanceof z.ZodError) {
      throw new ApiError(400, `Validation error: ${error.errors.map(e => e.message).join(', ')}`);
    }
    throw error;
  }
};

/**
 * @uri /api/v1/students/:id
 * @method PUT
 * @description Update a student
 * @param {{ body: StudentPayload, params: { id: string } }} req 
 */
const handleUpdateStudent = async (req, res) => {
  try {
    // Validate request
    const payload = z.object({
      userId: z.coerce.number().int("User ID must be a whole number"),
      name: z.string().min(3, "Name must be at least 3 characters").max(100, "Name must be less than 100 characters"),
      email: z.string().email("Invalid email format"),
      gender: z.enum(["Male", "Female", "Other"], { errorMap: () => ({ message: "Gender must be Male, Female, or Other" }) }),
      phone: z.string().length(10, "Phone number must be exactly 10 digits").regex(/^\d+$/, "Phone number must contain only digits"),
      dob: z.string(),
      admissionDate: z.string(),
      class: z.string().min(1, "Class is required").max(100, "Class name must be less than 100 characters"),
      section: z.string().min(1, "Section is required").max(100, "Section name must be less than 100 characters"),
      roll: z.coerce.number().int("Roll must be a whole number").min(1, "Roll must be greater than 0"),
      currentAddress: z.string().min(1, "Current address is required").max(100, "Current address must be less than 100 characters"),
      permanentAddress: z.string().min(1, "Permanent address is required").max(100, "Permanent address must be less than 100 characters"),
      fatherName: z.string().min(1, "Father's name is required").max(100, "Father's name must be less than 100 characters"),
      fatherPhone: z.string().length(10, "Father's phone must be exactly 10 digits").regex(/^\d+$/, "Father's phone must contain only digits"),
      motherName: z.string().min(1, "Mother's name is required").max(100, "Mother's name must be less than 100 characters"),
      motherPhone: z.string().length(10, "Mother's phone must be exactly 10 digits").regex(/^\d+$/, "Mother's phone must contain only digits"),
      guardianName: z.string().min(1, "Guardian's name is required").max(100, "Guardian's name must be less than 100 characters"),
      guardianPhone: z.string().length(10, "Guardian's phone must be exactly 10 digits").regex(/^\d+$/, "Guardian's phone must contain only digits"),
      relationOfGuardian: z.string().min(1, "Relation of guardian is required").max(100, "Relation of guardian must be less than 100 characters"),
      systemAccess: z.boolean().optional().default(false)
    }).parse({ ...req.body, userId: req.params.id });

    const message = await updateStudent(payload);
    res.json(message);
  } catch (error) {
    if (error instanceof z.ZodError) {
      throw new ApiError(400, `Validation error: ${error.errors.map(e => e.message).join(', ')}`);
    }
    throw error;
  }
};

/**
 * @uri /api/v1/students/:id
 * @method GET
 * @description Get a student detail
 * @param {{ params: { id: string } }} req
 */
const handleGetStudentDetail = async (req, res) => {
  try {
    const { id } = z.object({
      id: z.string().min(1, "Student ID is required")
    }).parse(req.params);
    
    const student = await getStudentDetail(id);
    res.json(student);
  } catch (error) {
    if (error instanceof z.ZodError) {
      throw new ApiError(400, `Validation error: ${error.errors.map(e => e.message).join(', ')}`);
    }
    throw error;
  }
};

/**
 * @uri /api/v1/students/:id/status
 * @method PUT
 * @description Update a student status
 * @param {{ body: { status: boolean }, params: { id: string }, user: { id: number } }} req 
 */
const handleStudentStatus = async (req, res) => {
  try {
    // Validate request
    const payload = z.object({
      status: z.boolean("Status must be a boolean value"),
      userId: z.coerce.number().int("User ID must be a whole number").min(1, "User ID must be greater than 0"),
      reviewerId: z.number().int("Reviewer ID must be a whole number").min(1, "Reviewer ID must be greater than 0"),
    }).parse({ ...req.body, userId: req.params.id, reviewerId: req.user.id });

    const message = await setStudentStatus(payload);
    res.json(message);
  } catch (error) {
    if (error instanceof z.ZodError) {
      throw new ApiError(400, `Validation error: ${error.errors.map(e => e.message).join(', ')}`);
    }
    throw error;
  }
};

module.exports = {
  handleGetAllStudents,
  handleGetStudentDetail,
  handleAddStudent,
  handleStudentStatus,
  handleUpdateStudent,
};
