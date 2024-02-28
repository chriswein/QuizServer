export interface QuestionItem {
    text: string,
    correct: boolean
}

export interface QuestionSet {
    text: string
    answers: QuestionItem[]
}

// let a: QuestionSet = {
//     text: "Wie heißt eine Spielekonsole?",
//     answers: [
//         { text: "Gintendo Gamegirl", correct: false },
//         { text: "Galapagos 5", correct: true },
//         { text: "Schaukausen", correct: false },
//         { text: "Playstation", correct: false },
//     ]
// }
import axios, { AxiosResponse } from 'axios';
import { PORT } from './config';

export class QuestionHandler {

    constructor() {
        axios.defaults.baseURL = `http://localhost:${PORT}`;
    }
    getNewQuestion(): Promise<QuestionSet> {
        return new Promise((resolve, _) => {
           
            axios.get<QuestionSet>('/question')
                .then((response: AxiosResponse<QuestionSet>) => {
                    console.log(response.data)
                    resolve(response.data)
                })
                .catch((error) => {
                    console.error('Error fetching data:', error);
                });

        });
    }
}