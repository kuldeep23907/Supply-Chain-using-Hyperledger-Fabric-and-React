const authRouter = require('express').Router();

authRouter.post('/users/:role', controller.signup);