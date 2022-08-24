// TODO: This is organized by frontend application, but it should be organized by business category

describe('ide', () => {
  describe('home page', () => {
    it('should return to home page when clicking icon', () => {
      // arrange
      cy.visit('http://localhost:3000/')
      
      // act
      cy.get('h1').click()
  
      // assert
      cy.get('h2').should('include.text', 'Ready to dive in? Your New Virtual BoH: NCX');
    })

    it('should navigate to ide when "get started" button clicked"', () => {
      // arrange
      cy.visit('http://localhost:3000/')

      // act
      cy.get('a').contains('Get Started').click()

      // assert
      cy.url() === 'http://localhost:3000/ide' // TODO: This is definitely wrong
      cy.get('h2').should('include.text', 'NCX')
      cy.get('a').should('include.text', 'Edit Entities')
      cy.get('a').should('include.text', 'Edit Business Rules')
    })
  })

  describe('entities page', () => {
    // placeholder
  })

  describe('rules page', () => {
    // placeholder
  })
})

describe('pos', () => {
  describe('item list', () => {
    it('shows item list header', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('h2').contains('Items List');
    })

    it('has 10 items', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('#item-list').children().should('have.length', 10)
    })
  })

  describe('total', () => {
    it('shows total header', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('h2').contains('Total');
    })

    it('has initial total of zero', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('#total-list').children().should('have.length', 0)
    })

    it('total list contains 1 element when item clicked', () => {
      // arrange
      cy.visit('http://localhost:3001');
      const item = cy.get('#item-list').children().first();

      // act
      item.click()

      // assert
      cy.get('#total-list').children().should('have.length', 1)
    })

    it('total list contains 2 elements when two items clicked', () => {
      // arrange
      cy.visit('http://localhost:3001');
      const firstItem = cy.get('#item-list').children().first();
      firstItem.click()
      const secondItem = firstItem.next();

      // act
      secondItem.click()

      // assert
      cy.get('#total-list').children().should('have.length', 2)
    })
  })

  describe('subtotal', () => {
    it('shows subtotal header', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('h3').contains('Subtotal');
    })

    it('has initial subtotal of zero', () => {
      // arrange
      cy.visit('http://localhost:3001');

      // assert
      cy.get('#subtotal').should('include.text', '$0.00')
    })

    it('subtotal $10.00 when purchasing 1 item', () => {
      // arrange
      cy.visit('http://localhost:3001');
      const item = cy.get('#item-list').children().first();

      // act
      item.click()

      // assert
      cy.get('#subtotal').should('include.text', '$10.00')
    })

    it('subtotal $20.00 when purchasing 2 items', () => {
      // arrange
      cy.visit('http://localhost:3001');
      const firstItem = cy.get('#item-list').children().first();
      firstItem.click()
      const secondItem = firstItem.next()

      // act
      secondItem.click()

      // assert
      cy.get('#subtotal').should('include.text', '$20.00')
    })
  })
})

