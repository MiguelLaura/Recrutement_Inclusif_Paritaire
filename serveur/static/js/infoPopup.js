class Message {

    static DEFAULT_TIMEOUT = 3; // en secondes

    // messages types

    static INFO = 0;
    static ERROR = 1;

    constructor(content, timeoutAfter = Message.DEFAULT_TIMEOUT) {
        this.content = content;
        this.timeoutTime = timeoutAfter;
        this.timeout = undefined;
        this.animation = undefined;
        this.parent = undefined;
        this.messageElement = undefined;
        this.loadingElt = undefined;
        this.closed = false;
    }

    createElement(parent, type) {
        const currObj = this;

        // Le conteneur du message
        const elt = document.createElement("div");
        elt.classList.add("infopopup-message", "hidden");

        // La barre de chargement
        const loading = document.createElement("p");
        loading.classList.add("infopopup-message-load")
        elt.appendChild(loading);

        // L'en-tête avec le bouton pour fermer

        const headerElt = document.createElement("header");
        const closeElt = document.createElement("i");
        closeElt.classList.add("bi", "bi-x", "clickable");
        closeElt.addEventListener("click", () => { currObj.close(); });
        headerElt.appendChild(closeElt);
        elt.appendChild(headerElt);

        // Le contenu du message

        const contentElt = document.createElement("p");
        contentElt.textContent = this.content;
        elt.appendChild(contentElt);

        // Insère le message en haut

        parent.insertAdjacentElement("afterbegin", elt);

        // Selon le type de message, le style va différer

        switch(type) {
            case Message.ERROR:
                elt.classList.add("err-msg");
            break
            default:
                elt.classList.add("info-msg");
        }

        this.messageElement = elt;
        this.loadingElt = loading;
        return this.messageElement;
    }

    // Anime la barre de chargement
    async animateLoading() {
        let animationFrame = [
            {
                "width": "0"
            }
        ];
    
        let animationTiming = {
            "duration" : this.timeoutTime * 1000,
            "fill": "forwards"
        };
    
        const animation = this.loadingElt.animate(animationFrame, animationTiming);
        this.animation = animation;

        try {
            await animation.finished
            animation.commitStyles();
            animation.cancel();
            this.close();
        } catch(e) {}
    }

    // Rend le message visible pendant un certain temps
    show() {
        if(!this.timeout) {
            const currObj = this;
            this.timeout = setTimeout(() => { currObj.close(); }, this.timeoutTime * 1000);
            this.animateLoading();
            this.messageElement.classList.remove("hidden");
        }
    }

    // Ferme le message
    close() {
        if(this.timeout !== undefined) {
            this.closed = true;
            clearTimeout(this.timeout);
            this.onclose();
            try {
                this.animation.commitStyles();
                this.animation.cancel();
            } catch(e) {}
            this.messageElement.remove();
            this.messageElement.classList.add("hidden");
        }
    }

    onclose() {}
}

class InfoPopup {

    // Le nombre de messages affichés en même temps
    static HISTORY_SIZE = 3;

    constructor() {
        this.messages = [];
        this.messageContainer = this.createContainer();
    }

    // Créer le conteneur de messages
    createContainer() {
        const container = document.createElement("div");
        container.classList.add("infopopup-container", "hidden");

        document.body.appendChild(container);

        return container;
    }

    // Ajoute le Message au conteneur
    addMessage(message) {
        if(this.messages.length == InfoPopup.HISTORY_SIZE) {
            const lastMsg = this.messages.pop();
            lastMsg.close();
        }

        this.messages.unshift(message);
        this.messageContainer.classList.remove("hidden");
        message.show();
    }

    _messageClosed() {
        if(this.messages.every((v => v.closed))) {
            this.messageContainer.classList.add("hidden");
        }
    }

    // Affiche un message d'info qui disparaîtra après un certain temps donné (en s)
    info(content, timeoutAfter = Message.DEFAULT_TIMEOUT) {
        const msg = new Message(content, timeoutAfter, 0);
        msg.onclose = () => { this._messageClosed(); }
        msg.createElement(this.messageContainer, Message.INFO);
        this.addMessage(msg);
    }

    // Affiche un message d'erreur qui disparaîtra après un certain temps donné (en s)
    error(content, timeoutAfter = Message.DEFAULT_TIMEOUT) {
        const msg = new Message(content, timeoutAfter, 0);
        msg.onclose = () => { this._messageClosed(); }
        msg.createElement(this.messageContainer, Message.ERROR);
        this.addMessage(msg);
    }
}