const asyncHandler = require("express-async-handler");
const express = require("express");
const router = express.Router();
const studentController = require("./students-controller");

router.get("", asyncHandler(studentController.handleGetAllStudents));
router.post("", asyncHandler(studentController.handleAddStudent));
router.get("/:id", asyncHandler(studentController.handleGetStudentDetail));
router.post("/:id/status", asyncHandler(studentController.handleStudentStatus));
router.put("/:id", asyncHandler(studentController.handleUpdateStudent));

module.exports = { studentsRoutes: router };
