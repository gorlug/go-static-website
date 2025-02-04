const { selectWebsiteLanguage } = require('../language-select');

describe('Language Detection', () => {
    const originalLanguage = navigator.language;
    const langs = {
        de: '/de.html',
        en: '/en.html',
    }
    let mockGetItem, mockSetItem, mockReplace;

    beforeEach(() => {
        const { getItem, setItem } = mockLocalStorage();
        mockGetItem = getItem;
        mockSetItem = setItem;
        mockReplace = mockWindowLocationReplace();
    });

    function mockNavigatorLanguage(language) {
        Object.defineProperty(navigator, 'language', {
            configurable: true,
            get: () => language
        });
    }

    afterEach(() => {
        // Restore original
        Object.defineProperty(navigator, 'language', {
            configurable: true,
            get: () => originalLanguage
        });
    });

    it('should select the browser language url', () => {
        mockNavigatorLanguage('de-DE');
        mockGetItem.mockReturnValue(null);

        selectWebsiteLanguage('en', langs, 'en');

        expect(mockGetItem).toHaveBeenCalledWith('initialLangSelected');
        expect(mockSetItem).toHaveBeenCalledWith('initialLangSelected', 'true');
        expect(mockReplace).toHaveBeenCalledWith('/de.html');
    });

    it('should select the default language if it is not in langs', () => {
        mockNavigatorLanguage('nl-NL');
        mockGetItem.mockReturnValue(null);

        selectWebsiteLanguage('de', langs, 'en');

        expect(mockGetItem).toHaveBeenCalledWith('initialLangSelected');
        expect(mockSetItem).toHaveBeenCalledWith('initialLangSelected', 'true');
        expect(mockReplace).toHaveBeenCalledWith('/en.html');
    })

    it('should not redirect if the browser language is the same as the current language', () => {
        mockNavigatorLanguage('en-US');
        mockGetItem.mockReturnValue(null);

        selectWebsiteLanguage('en', langs, 'en');

        expect(mockGetItem).toHaveBeenCalledWith('initialLangSelected');
        expect(mockSetItem).toHaveBeenCalledWith('initialLangSelected', 'true');
        expect(mockReplace).not.toHaveBeenCalled();
    })

    it('should not redirect if the language was already selected', () => {
        mockNavigatorLanguage('de-DE');
        mockGetItem.mockReturnValue('true');

        selectWebsiteLanguage('en', langs, 'en');

        expect(mockGetItem).toHaveBeenCalledWith('initialLangSelected');
        expect(mockSetItem).not.toHaveBeenCalled();
        expect(mockReplace).not.toHaveBeenCalled();
    })
});

function mockLocalStorage() {
    const getItem = jest.fn();
    const setItem = jest.fn();

    Object.defineProperty(window, 'localStorage', {
        value: {
            getItem,
            setItem
        }
    });
    return { getItem, setItem };
}

function mockWindowLocationReplace() {
    const replace = jest.fn();
    Object.defineProperty(window, 'location', {
        value: {
            replace
        }
    });
    return replace;
}
