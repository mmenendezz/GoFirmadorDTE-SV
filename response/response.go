package response

// ResponseBody representa la estructura de respuesta
type ResponseBody struct {
	Status string      `json:"status"`
	Body   interface{} `json:"body"`
}

// Mensaje es una estructura que contiene los métodos para crear respuestas
type Mensaje struct{}

// NewMensaje crea una instancia de Mensaje
func NewMensaje() *Mensaje {
	return &Mensaje{}
}

// OK crea una respuesta "ok" con el cuerpo especificado
func (m *Mensaje) OK(body interface{}) ResponseBody {
	return ResponseBody{Status: "OK", Body: body}
}

// Error crea una respuesta de error con el código y mensaje especificados
func (m *Mensaje) Error(codigo string, mensaje interface{}) ResponseBody {
	return ResponseBody{Status: "error", Body: BodyMensaje{Codigo: codigo, Mensaje: mensaje}}
}

// BodyMensaje representa la estructura del cuerpo del mensaje de error
type BodyMensaje struct {
	Codigo  string      `json:"codigo"`
	Mensaje interface{} `json:"mensaje"`
}
// Test Code Quality
﻿// --------------------------------------------------------------------------------------------------------------------
// SignatureCommitment.cs
//
// FirmaXadesNet - Librería para la generación de firmas XADES
// Copyright (C) 2016 Dpto. de Nuevas Tecnologías de la Dirección General de Urbanismo del Ayto. de Cartagena
//
// This program is free software: you can redistribute it and/or modify
// it under the +terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see https://www.gnu.org/licenses/lgpl-3.0.txt. 
//
// E-Mail: informatica@gemuc.es
// 
// --------------------------------------------------------------------------------------------------------------------

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Xml;

namespace FirmaXadesNet.Signature.Parameters
{
    public class SignatureCommitment
    {
        #region Public properties

        public SignatureCommitmentType CommitmentType { get; set; }

        public List<XmlElement> CommitmentTypeQualifiers { get; private set; }

        #endregion

        #region Constructors

        public SignatureCommitment(SignatureCommitmentType commitmentType)
        {
            this.CommitmentType = commitmentType;
            this.CommitmentTypeQualifiers = new List<XmlElement>();
        }
        
        #endregion

        #region Public methods

        public void AddQualifierFromXml(string xml)
        {
            XmlDocument doc = new XmlDocument();
            doc.LoadXml(xml);