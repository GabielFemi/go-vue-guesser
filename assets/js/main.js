const guessComponent = {
    template:
    `<input name="guess" class="input is-64x64" type="text" v-model = "input" @keyup.enter = "monitorEnterKey"/>`,
    data () {
        return {
            input: ''
        }
    }
}
new Vue({
    el: '#app',
    data: {
        notes: [],
        guesses: [],
        placeholder: 'Guess what I\'m thinking'
    },
    computed: {

    },
    components: {
      'guess-component': guessComponent
    }
})