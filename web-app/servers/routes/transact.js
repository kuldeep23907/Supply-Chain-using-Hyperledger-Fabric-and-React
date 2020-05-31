const transactRouter = require('express').Router();
const controller = require('../controllers/transact.js');
const authMiddleware = require('../middlewares/auth.js');

transactRouter.use('/', authMiddleware);

transactRouter.post('/', controller.transactProduct);

transactRouter.use('/consumer', authMiddleware);

transactRouter.post('/consumer', controller.transactProductConsumer);
module.exports = transactRouter;
