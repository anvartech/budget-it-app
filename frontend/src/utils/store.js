import {
    createStore,
    applyMiddleware,
    combineReducers,
} from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension/logOnlyInProduction';

import createSagaMiddleware from 'redux-saga';

const sagaMiddleware = createSagaMiddleware();

export const reducer = combineReducers({});

export const store = createStore(
    reducer,
    composeWithDevTools(
        applyMiddleware(...sagaMiddleware),
    )
);

sagaMiddleware.run();
