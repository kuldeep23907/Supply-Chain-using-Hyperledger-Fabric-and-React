const transactRouter = require('express').Router();
const controller = require('../controllers/transact.js');
const authMiddleware = require('../middlewares/auth.js');

transactRouter.use('/transact', authMiddleware);

transactRouter.post('/transact', controller.transactProduct);