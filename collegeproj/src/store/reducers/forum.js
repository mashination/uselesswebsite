import { PUBLISH_TOPIC, GET_TOPIC, GET_TOPICS, PUBLISH_REPLY , CLEAR_TOPIC, GET_LAUNCHES} from "../actions/forum";

const initialState = {
 username : null,
 currentTopic : null,
 topics : null,
 launches : null
}

export default (state = initialState, action) => {
    switch (action.type) {
        case PUBLISH_TOPIC :
            return state;
        case PUBLISH_REPLY :
            return state;
        case GET_TOPIC : 
            return {
                ...state,
                currentTopic : action.topic

            }
        case GET_TOPICS : 
            return {
                ...state,
                topics : action.topics

            }
        case CLEAR_TOPIC :
            return{
                ...state,
                currentTopic : null
            }
        case GET_LAUNCHES :
            return{
                ...state,
                launches : action.launches
            }

        default :
            return state;
    }
}

