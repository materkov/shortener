import React from 'react';
import style from './shortener.module.css';
import axios from "axios";

class Shortener extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    handleKeyDown = (event) => {
        if (event.key === 'Enter') {
            this.apiRequest(this.state.url);
        }
    };

    handleFocus = (event) => {
        event.target.select();
    };

    handleChange = (event) => {
        this.setState({...this.state, url: event.target.value});
    };

    async apiRequest(url) {
        this.setState({...this.state, error: null, shortenedUrl: null});

        let apiHost = "https://shortener.mmaks.me";
        if (window.location.host === "localhost:8000") {
            apiHost = "http://localhost:8000";
        }

        let redirectHost = "https://s.mmaks.me";
        if (window.location.host === "localhost:8000") {
            redirectHost = "http://localhost:8000/redirect";
        }

        try {
            const response = await axios.post(apiHost + "/create", {url: url});

            let shortenedUrl = redirectHost + "/" + response.data.key;
            this.setState({...this.state, error: null, shortenedUrl: shortenedUrl});
        } catch (error) {
            let err = "Кажется, что-то пошло не так";
            if (error.response.data.error === "bad_url") {
                err = "Некорректный URL";
            }

            this.setState({...this.state, error: err})
        }
    }

    render() {
        let shortenedUrl;

        if (this.state.shortenedUrl) {
            shortenedUrl =
                <div>
                    <div className={style.top_label}>Сокращенный URL:</div>
                    <input type="text" className={style.url_input} value={this.state.shortenedUrl} autoFocus
                           readOnly={true}
                           onFocus={this.handleFocus}/>
                </div>
        }

        return (
            <div className={style.container}>
                <div className={style.top_label}>Введите тут свой URL:</div>
                <input type="text" className={style.url_input} onKeyDown={this.handleKeyDown}
                       onChange={this.handleChange}/>

                <div className={style.error_label}>{this.state.error}</div>

                {shortenedUrl}
            </div>
        )
    }
}

export default Shortener;
