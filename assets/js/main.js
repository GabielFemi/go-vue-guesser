const guessComponent = {
    template:
    `<input name="guess" class="input is-64x64" type="text" v-model = "input" @keyup.enter = "monitorEnterKey"/>`,
    data () {
        return {
            input: ''
        }
    },
    methods: {
        monitorEnterKey() {
            this.$emit('add-note', {
                note: this.input,
            });
            this.input = '';
        },
        addNote(event) {
            this.notes.push(event.note);
        }
    },
};

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