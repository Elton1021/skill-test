const { ApiError } = require("../utils");

const handleGlobalError = (err, req, res, next) => {
    console.error(err);
    // Log error details with process.stdout.write as console.log is not working
    process.stdout.write(JSON.stringify({ 
        message: err.message, 
        stack: err.stack, 
        statusCode: err.code || err.statusCode 
    }) + '\n');
    
    if (err instanceof ApiError) {
        return res.status(err.statusCode).json({ error: err.message });
    }

    return res.status(500).json({ error: "Internal server error" });
}

module.exports = { handleGlobalError };
