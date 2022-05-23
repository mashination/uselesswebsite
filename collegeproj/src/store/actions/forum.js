
export const PUBLISH_TOPIC = 'PUBLISH_TOPIC'
export const PUBLISH_REPLY = 'PUBLISH_REPLY'
export const GET_TOPIC = 'GET_TOPIC'
export const GET_TOPICS = 'GET_TOPICS'
export const GET_LAUNCHES = 'GET_LAUNCHES' 

export const CLEAR_TOPIC = 'CLEAR_TOPIC' 

export const publishTopic = (username, title, content) => {
    
    console.log(username)
        console.log(title)
        console.log(content)
    return async dispatch => {
        console.log(username)
        console.log(title)
        console.log(content)
        const response = await fetch( 'http://localhost:10000/topic', {
            method : 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body : JSON.stringify({
                usr : username,
                title : title,
                content: content
            })
        })
        if (!response.ok){
            const resData = await response.json();
            console.log(response.status)
        }
        const resData = await response.json();
        dispatch({type: PUBLISH_TOPIC})
    }
}
export const publishReply = (username, topicid, content) => {
    
    return async dispatch => {
        console.log(username)
        console.log(topicid)
        console.log(content)
        const response = await fetch( 'http://localhost:10000/reply', {
            method : 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body : JSON.stringify({
                usr : username,
                topicid : topicid,
                content: content
            })
        })
        if (!response.ok){
            const resData = await response.json();
            console.log(response.status)
        }
        const resData = await response.json();
        dispatch({type: PUBLISH_REPLY})
    }
}

export const getTopic = (topicid) => {
    
    return async dispatch => {
        
        const response = await fetch( `http://localhost:10000/topic/${topicid}`, {
            method : 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            
        })
        if (!response.ok){
            const resData = await response.json();
            console.log(response.status)
        }
        const resData = await response.json();
        console.log(resData)
        dispatch({type: GET_TOPIC, topic : resData})
    }
}

export const getTopics = () => {
    
    return async dispatch => {
        
        const response = await fetch( `http://localhost:10000/topics`, {
            method : 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            
        })
        if (!response.ok){
            const resData = await response.json();
            console.log(response.status)
        }
        const resData = await response.json();
        console.log(resData)
        dispatch({type: GET_TOPICS, topics : resData})
    }
}
export const getLaunches = () => {
    
    return async dispatch => {
        
        const response = await fetch( `http://localhost:10000/launches`, {
            method : 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            
        })
        if (!response.ok){
            const resData = await response.json();
            console.log(response.status)
        }
        const resData = await response.json();
        console.log(resData)
        dispatch({type: GET_LAUNCHES, launches : resData})
    }
}

export const clearCurrentTopic = () => {
    return {type : CLEAR_TOPIC}
}