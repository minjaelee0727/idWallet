//
//  NetworkService.swift
//  idWallet
//
//  Created by Min Jae Lee on 2022-05-29.
//

import Foundation
import Combine

class NetworkService: ObservableObject {
    private var cancellable = Set<AnyCancellable>()
    
    func getWallet(rq: WalletRequest) {
        guard let url = Constants.registerURL else {
            fatalError("InvalidURL")
        }
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.httpBody = try? JSONSerialization.data(withJSONObject: ["birthDate": rq.Birthdate, "sin": rq.Sin])
        
        URLSession.shared
            .dataTaskPublisher(for: request)
            .map { $0.data }
            .decode(type: IdWallet.self, decoder: JSONDecoder())
            .eraseToAnyPublisher()
            .sink(receiveCompletion: {
                completion in
                
                switch completion {
                case .finished:
                    break
                case .failure(let error):
                    print(error.localizedDescription)
                }
                
            }, receiveValue: { data in
                UserDefaults.standard.set(data.publicKey, forKey: Constants.publicKey)
                // TODO: Save into Keychain
                UserDefaults.standard.set(data.secretKey, forKey: Constants.secretKey)
                print(data)
            })
            .store(in: &self.cancellable)
    }
}
